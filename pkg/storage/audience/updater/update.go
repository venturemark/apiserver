package updater

import (
	"encoding/json"
	"fmt"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Update(req *audience.UpdateI) (*audience.UpdateO, error) {
	var err error

	var aui float64
	{
		aui, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.AudienceID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cur []byte
	{
		k := fmt.Sprintf(key.Audience)
		s, err := u.redigo.Sorted().Search().Score(k, aui, aui)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		aud := &schema.Audience{}
		err = json.Unmarshal([]byte(s[0]), aud)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(aud)
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

	var aud schema.Audience
	{
		err := json.Unmarshal([]byte(val), &aud)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd bool
	{
		k := fmt.Sprintf(key.Audience)
		v := val
		s := aui
		i := index.New(index.Name, aud.Obj.Property.Name)

		upd, err = u.redigo.Sorted().Update().Value(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.UpdateO
	{
		res = &audience.UpdateO{
			Obj: []*audience.UpdateO_Obj{
				{
					Metadata: map[string]string{},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.AudienceStatus] = "updated"
		}
	}

	return res, nil
}
