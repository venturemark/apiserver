package auth

import (
	"context"
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
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

func (v *Verifier) Verify(ctx context.Context, req *venture.SearchI) (bool, error) {
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
		vei := met[metadata.VentureID]
		if vei == "" {
			return action.Filter, nil
		}
	}

	return action.Search, nil
}

func (v *Verifier) res(met map[string]string) (label.Label, error) {
	return resource.Venture, nil
}

func (v *Verifier) rol(met map[string]string) (label.Label, error) {
	var err error

	{
		vei := met[metadata.VentureID]
		if vei == "" {
			return role.Subject, nil
		}
	}

	var tim string
	{
		tim, err = v.resolveTimelineRole(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var ven string
	if tim != role.Any.Label() {
		ven, err = v.permission.Resolver().Venture().Role(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var rol label.Label
	{
		if tim == role.Any.Label() {
			rol = role.Any
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
	if met[metadata.UserID] == "" {
		return visibility.Public, nil
	}

	var err error
	var ven string
	{
		ven, err = v.permission.Resolver().Venture().Visibility(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var tim string
	{
		tim, err = v.resolveTimelineVisibility(met)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var vis label.Label
	{
		if tim == visibility.Public.Label() {
			vis = visibility.Public
		} else if ven == "" {
			vis = visibility.Private
		} else if ven == visibility.Private.Label() {
			vis = visibility.Private
		} else if ven == visibility.Member.Label() {
			vis = visibility.Member
		} else if ven == visibility.Public.Label() {
			vis = visibility.Public
		}
	}

	return vis, nil
}

func (v *Verifier) resolveTimelineRole(met map[string]string) (string, error) {
	if met[metadata.UserID] == "" {
		return role.Any.Label(), nil
	}

	var err error
	var tik *key.Key
	{
		tik = key.Timeline(met)
	}

	var str []string
	{
		k := tik.List()

		str, err = v.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var rol string
	{
		for _, s := range str {
			tim := &schema.Timeline{}
			err := json.Unmarshal([]byte(s), tim)
			if err != nil {
				return "", tracer.Mask(err)
			}

			if tim.Obj.Metadata[metadata.ResourceVisibility] == "public" {
				rol = role.Any.Label()
				break
			}

			rol, err = v.permission.Resolver().Timeline().Role(timeline(tim, met))
			if err != nil {
				return "", tracer.Mask(err)
			}

			if rol != "" {
				break
			}
		}
	}

	return rol, nil
}

func (v *Verifier) resolveTimelineVisibility(met map[string]string) (string, error) {
	var err error

	var tik *key.Key
	{
		tik = key.Timeline(met)
	}

	var str []string
	{
		k := tik.List()

		str, err = v.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var vis string
	{
		for _, s := range str {
			tim := &schema.Timeline{}
			err := json.Unmarshal([]byte(s), tim)
			if err != nil {
				return "", tracer.Mask(err)
			}

			if tim.Obj.Metadata[metadata.ResourceVisibility] == "public" {
				vis = visibility.Public.Label()
				break
			}

			vis = visibility.Private.Label()
		}
	}

	return vis, nil
}

func timeline(tim *schema.Timeline, met map[string]string) map[string]string {
	cop := map[string]string{}

	for k, v := range met {
		cop[k] = v
	}

	cop[metadata.ResourceKind] = "timeline"
	cop[metadata.TimelineID] = tim.Obj.Metadata[metadata.TimelineID]

	return cop
}
