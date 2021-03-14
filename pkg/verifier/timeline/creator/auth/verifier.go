package auth

import (
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

func (v *Verifier) Verify(req *timeline.CreateI) (bool, error) {
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
		vis, err = v.vis(req.Obj[0].Metadata)
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
	return action.Create, nil
}

func (v *Verifier) res(met map[string]string) (label.Label, error) {
	return resource.Timeline, nil
}

func (v *Verifier) rol(met map[string]string) (label.Label, error) {
	var err error

	var ven string
	{
		ven, err = v.permission.Resource().Venture().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var rol label.Label
	{
		if ven == role.Member.Label() {
			rol = role.Member
		}
		if ven == role.Owner.Label() {
			rol = role.Owner
		}
	}

	return rol, nil
}

func (v *Verifier) vis(met map[string]string) (label.Label, error) {
	var vis label.Label
	{
		mes := met[metadata.ResourceVisibility]

		if mes == "" {
			vis = visibility.Any
		}
		if mes == visibility.Any.Label() {
			vis = visibility.Any
		}
		if mes == visibility.Private.Label() {
			vis = visibility.Private
		}
		if mes == visibility.Public.Label() {
			vis = visibility.Public
		}
	}

	return vis, nil
}
