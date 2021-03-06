package updater

import (
	"encoding/json"
	"fmt"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Update(req *role.UpdateI) (*role.UpdateO, error) {
	var err error

	var rei string
	{
		rei = req.Obj[0].Metadata[metadata.ResourceID]
	}

	var roi float64
	{
		roi, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.RoleID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cur []byte
	{
		k := fmt.Sprintf(key.Role, rei)
		s, err := u.redigo.Sorted().Search().Score(k, roi, roi)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		rol := &schema.Role{}
		err = json.Unmarshal([]byte(s[0]), rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(rol)
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

	var upd bool
	{
		k := fmt.Sprintf(key.Role, rei)
		v := val
		s := roi

		upd, err = u.redigo.Sorted().Update().Value(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *role.UpdateO
	{
		res = &role.UpdateO{
			Obj: []*role.UpdateO_Obj{
				{
					Metadata: map[string]string{},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.RoleStatus] = "updated"
		}
	}

	return res, nil
}
