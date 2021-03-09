package updater

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

// Update provides a storage primitive to modify timelines associated with a
// user.
func (u *Updater) Update(req *timeline.UpdateI) (*timeline.UpdateO, error) {
	var err error

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	var cur []byte
	{
		k := tik.List()
		s, err := u.redigo.Sorted().Search().Score(k, tik.ID().F(), tik.ID().F())
		if err != nil {
			return nil, tracer.Mask(err)
		}

		tim := &schema.Timeline{}
		err = json.Unmarshal([]byte(s[0]), tim)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(tim)
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

	var tim schema.Timeline
	{
		err := json.Unmarshal([]byte(val), &tim)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd bool
	{
		k := tik.List()
		v := val
		s := tik.ID().F()
		i := index.New(index.Name, tim.Obj.Property.Name)

		upd, err = u.redigo.Sorted().Update().Value(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.UpdateO
	{
		res = &timeline.UpdateO{
			Obj: []*timeline.UpdateO_Obj{
				{
					Metadata: map[string]string{},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.TimelineStatus] = "updated"
		}
	}

	return res, nil
}
