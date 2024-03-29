package auth

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/label"
	"github.com/venturemark/permission/pkg/label/action"
	"github.com/venturemark/permission/pkg/label/resource"
	"github.com/venturemark/permission/pkg/label/role"
	"github.com/venturemark/permission/pkg/label/visibility"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type VerifierConfig struct {
	Permission permission.Gateway
	Redigo     redigo.Interface
}

type Verifier struct {
	permission permission.Gateway
	redigo     redigo.Interface
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	if config.Permission == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Permission must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	v := &Verifier{
		permission: config.Permission,
		redigo:     config.Redigo,
	}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *message.SearchI) (bool, error) {
	var err error

	var act label.Label
	var res label.Label
	var rol label.Label
	var vis label.Label
	{
		act, err = v.act(req.Obj[0].Metadata)
		if err != nil {
			return false, tracer.Mask(err)
		}
		res, err = v.res(req.Obj[0].Metadata)
		if err != nil {
			return false, tracer.Mask(err)
		}
		rol, err = v.rol(req.Obj[0].Metadata)
		if err != nil {
			return false, tracer.Mask(err)
		}
		vis, err = v.vis(ctx, req.Obj[0].Metadata)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	{
		ok := v.permission.Ingress().Match(act, res, rol, vis)
		if !ok {
			return false, nil
		}
	}

	return true, nil
}

func (v *Verifier) act(met map[string]string) (label.Label, error) {
	{
		mei := met[metadata.MessageID]
		if mei == "" {
			return action.Filter, nil
		}
	}

	return action.Search, nil
}

func (v *Verifier) res(met map[string]string) (label.Label, error) {
	return resource.Message, nil
}

func (v *Verifier) rol(met map[string]string) (label.Label, error) {
	var err error

	{
		mei := met[metadata.MessageID]
		if mei == "" {
			return role.Subject, nil
		}
	}

	var mes string
	{
		mes, err = v.permission.Resolver().Message().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var tim string
	{
		tim, err = v.permission.Resolver().Timeline().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var rol label.Label
	{
		if mes == role.Member.Label() {
			rol = role.Member
		}
		if mes == role.Owner.Label() {
			rol = role.Owner
		}
		if mes == role.Reader.Label() {
			rol = role.Reader
		}
		if tim == role.Member.Label() {
			rol = role.Member
		}
		if tim == role.Owner.Label() {
			rol = role.Owner
		}
		if tim == role.Reader.Label() {
			rol = role.Reader
		}
	}

	return rol, nil
}

func (v *Verifier) vis(ctx context.Context, met map[string]string) (label.Label, error) {
	var err error
	var tim string
	{
		tim, err = v.permission.Resolver().Timeline().Visibility(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var vis label.Label
	{
		if tim == "" {
			vis = visibility.Private
		}
		if tim == visibility.Private.Label() {
			vis = visibility.Private
		}
		if tim == visibility.Member.Label() {
			vis = visibility.Private
		}
		if tim == visibility.Public.Label() {
			vis = visibility.Public
		}
	}

	return vis, nil
}
