package updater

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"
)

// Update provides a storage primitive to modify text updates associated with a
// timeline. A timeline refers to many updates. For more information about
// technical details see the inline documentation.
func (u *Updater) Update(req *texupd.UpdateI) (*texupd.UpdateO, error) {
	var err error

	var upk *key.Key
	{
		upk = key.Update(req.Obj[0].Metadata)
	}

	var cur []byte
	{
		k := upk.List()
		s, err := u.redigo.Sorted().Search().Score(k, upk.ID().F(), upk.ID().F())
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
		k := upk.List()
		v := val
		s := upk.ID().F()

		upd, err = u.redigo.Sorted().Update().Value(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.UpdateO
	{
		res = &texupd.UpdateO{
			Obj: []*texupd.UpdateO_Obj{
				{
					Metadata: map[string]string{},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.UpdateStatus] = "updated"
		}
	}

	return res, nil
}
