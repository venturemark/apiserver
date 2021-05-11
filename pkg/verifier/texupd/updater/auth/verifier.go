package auth

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
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

func (v *Verifier) Verify(ctx context.Context, req *texupd.UpdateI) (bool, error) {
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
	return action.Update, nil
}

func (v *Verifier) res(met map[string]string) (label.Label, error) {
	return resource.Update, nil
}

func (v *Verifier) rol(met map[string]string) (label.Label, error) {
	var err error

	var tim string
	{
		tim, err = v.permission.Resolver().Timeline().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var upd string
	{
		upd, err = v.permission.Resolver().Update().Role(met)
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
		if tim == role.Owner.Label() {
			rol = role.Owner
		}
		if upd == role.Owner.Label() {
			rol = role.Owner
		}
		if ven == role.Owner.Label() {
			rol = role.Owner
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

	var upd string
	{
		upd, err = v.permission.Resolver().Update().Visibility(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var vis label.Label
	{
		if upd == "" {
			vis = visibility.Any
		}
		if upd == visibility.Any.Label() {
			vis = visibility.Any
		}
		if upd == visibility.Private.Label() {
			vis = visibility.Private
		}
		if upd == visibility.Public.Label() && isp {
			vis = visibility.Public
		}
	}

	return vis, nil
}
