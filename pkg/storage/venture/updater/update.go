package updater

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Update(req *venture.UpdateI) (*venture.UpdateO, error) {
	var err error

	var vek *key.Key
	{
		vek = key.Venture(req.Obj[0].Metadata)
	}

	var cur []byte
	{
		k := vek.Elem()

		str, err := u.redigo.Simple().Search().Value(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		ven := &schema.Venture{}
		err = json.Unmarshal([]byte(str), ven)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(ven)
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

	var ven schema.Venture
	{
		err := json.Unmarshal([]byte(val), &ven)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd bool
	{
		k := vek.Elem()
		v := val

		err = u.redigo.Simple().Create().Element(k, v)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		// For now we assume the venture got updated.
		upd = true
	}

	var res *venture.UpdateO
	{
		res = &venture.UpdateO{
			Obj: []*venture.UpdateO_Obj{
				{
					Metadata: map[string]string{
						metadata.VentureID: vek.ID().S(),
					},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.VentureStatus] = "updated"
		}
	}

	return res, nil
}
