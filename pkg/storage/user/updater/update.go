package updater

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Update(req *user.UpdateI) (*user.UpdateO, error) {
	var err error

	var usk *key.Key
	{
		usk = key.User(req.Obj[0].Metadata)
	}

	var cur []byte
	{
		k := usk.Elem()

		str, err := u.redigo.Simple().Search().Value(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		use := &schema.User{}
		err = json.Unmarshal([]byte(str), use)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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

	var use schema.User
	{
		err := json.Unmarshal([]byte(val), &use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd bool
	{
		k := usk.Elem()
		v := val

		err = u.redigo.Simple().Create().Element(k, v)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		// For now we assume the user got updated.
		upd = true
	}

	var res *user.UpdateO
	{
		res = &user.UpdateO{
			Obj: []*user.UpdateO_Obj{
				{
					Metadata: map[string]string{
						metadata.UserID: usk.ID().S(),
					},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.UserStatus] = "updated"
		}
	}

	return res, nil
}
