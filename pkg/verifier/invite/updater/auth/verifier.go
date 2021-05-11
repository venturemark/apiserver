package auth

import (
	"context"
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
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

func (v *Verifier) Verify(ctx context.Context, req *invite.UpdateI) (bool, error) {
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
	return resource.Invite, nil
}

func (v *Verifier) rol(met map[string]string) (label.Label, error) {
	var err error

	var cur string
	var des string
	{
		cur = met[metadata.InviteCode]

		des, err = v.inviteCode(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var inv string
	{
		inv, err = v.permission.Resolver().Invite().Role(met)
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
		if cur == des {
			rol = role.Owner
		}
		if inv == role.Owner.Label() {
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

	var inv string
	{
		inv, err = v.permission.Resolver().Invite().Visibility(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var vis label.Label
	{
		if inv == "" {
			vis = visibility.Any
		}
		if inv == visibility.Any.Label() {
			vis = visibility.Any
		}
		if inv == visibility.Private.Label() {
			vis = visibility.Private
		}
		if inv == visibility.Public.Label() && isp {
			vis = visibility.Public
		}
	}

	return vis, nil
}

func (v *Verifier) inviteCode(met map[string]string) (string, error) {
	var ink *key.Key
	{
		ink = key.Invite(met)
	}

	var inc string
	{
		k := ink.List()
		s := ink.ID().F()

		str, err := v.redigo.Sorted().Search().Score(k, s, s)
		if err != nil {
			return "", tracer.Mask(err)
		}

		if len(str) != 1 {
			return "", nil
		}

		inv := &schema.Invite{}
		err = json.Unmarshal([]byte(str[0]), inv)
		if err != nil {
			return "", tracer.Mask(err)
		}

		inc = inv.Obj.Metadata[metadata.InviteCode]
	}

	return inc, nil
}
