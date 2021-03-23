package updater

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Update(req *invite.UpdateI) (*invite.UpdateO, error) {
	var err error

	var ink *key.Key
	{
		ink = key.Invite(req.Obj[0].Metadata)
	}

	var cur []byte
	var old string
	{
		k := ink.List()
		s := ink.ID().F()

		str, err := u.redigo.Sorted().Search().Score(k, s, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		inv := &schema.Invite{}
		err = json.Unmarshal([]byte(str[0]), inv)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(inv)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		old = inv.Obj.Property.Stat
	}

	var pat []byte
	{
		var p []map[string]string

		for _, j := range req.Obj[0].Jsnpatch {
			m := map[string]string{
				"op":    j.GetOpe(),
				"path":  j.GetPat(),
				"value": j.GetVal(),
			}

			p = append(p, m)
		}

		pat, err = json.Marshal(p)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var val string
	{
		patch, err := jsonpatch.DecodePatch(pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		des, err := patch.Apply(cur)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(des)
	}

	var inv schema.Invite
	var new string
	{
		err := json.Unmarshal([]byte(val), &inv)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		new = inv.Obj.Property.Stat
	}

	var acc bool
	{
		if old == "pending" && new == "accepted" {
			acc = true
		}
	}

	var upd bool
	{
		k := ink.List()
		v := val
		s := ink.ID().F()
		i := index.New(index.Mail, inv.Obj.Property.Mail)

		upd, err = u.redigo.Sorted().Update().Value(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *invite.UpdateO
	{
		res = &invite.UpdateO{
			Obj: []*invite.UpdateO_Obj{
				{
					Metadata: map[string]string{},
				},
			},
		}

		if acc {
			res.Obj[0].Metadata[metadata.RoleStatus] = "created"
		}

		if upd {
			res.Obj[0].Metadata[metadata.InviteStatus] = "updated"
		}
	}

	return res, nil
}
