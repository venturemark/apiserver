package auth

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/label"
	"github.com/venturemark/permission/pkg/label/action"
	"github.com/venturemark/permission/pkg/label/resource"
	"github.com/venturemark/permission/pkg/label/role"
	"github.com/venturemark/permission/pkg/label/visibility"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/claimid"
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

func (v *Verifier) Verify(ctx context.Context, req *timeline.SearchI) (bool, error) {
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
		tii := met[metadata.TimelineID]
		if tii == "" {
			return action.Filter, nil
		}
	}

	return action.Search, nil
}

func (v *Verifier) res(met map[string]string) (label.Label, error) {
	return resource.Timeline, nil
}

func (v *Verifier) rol(met map[string]string) (label.Label, error) {
	var err error

	{
		tii := met[metadata.TimelineID]
		if tii == "" {
			return role.Subject, nil
		}
	}

	var tim string
	{
		tim, err = v.permission.Resolver().Timeline().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var ven string
	{
		ven, err = v.permission.Resolver().Venture().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var rol label.Label
	{
		if tim == role.Member.Label() {
			rol = role.Member
		}
		if tim == role.Owner.Label() {
			rol = role.Owner
		}
		if tim == role.Reader.Label() {
			rol = role.Reader
		}
		if ven == role.Member.Label() {
			rol = role.Member
		}
		if ven == role.Owner.Label() {
			rol = role.Owner
		}
		if ven == role.Reader.Label() {
			rol = role.Reader
		}
	}

	return rol, nil
}

func (v *Verifier) vis(ctx context.Context, met map[string]string) (label.Label, error) {
	var err error

	var isp bool
	{
		cli, _ := claimid.FromContext(ctx)
		if cli == "webclient" {
			isp = true
		}
	}

	var ven string
	{
		ven, err = v.permission.Resolver().Venture().Visibility(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var vis label.Label
	{
		if ven == "" {
			vis = visibility.Private
		}
		if ven == visibility.Private.Label() {
			vis = visibility.Private
		}
		if ven == visibility.Public.Label() && isp {
			vis = visibility.Public
		}
	}

	return vis, nil
}
