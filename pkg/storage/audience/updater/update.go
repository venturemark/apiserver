package updater

import (
	"encoding/json"
	"fmt"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/index"
	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/audience/element"
)

func (u *Updater) Update(req *audience.UpdateI) (*audience.UpdateO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var aid float64
	{
		aid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.AudienceID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cur []byte
	{
		k := fmt.Sprintf(key.Audience, oid)
		str, err := u.redigo.Sorted().Search().Score(k, aid, aid)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		_, n, t, u, err := element.Split(str[0])
		if err != nil {
			return nil, tracer.Mask(err)
		}

		obj := &audience.CreateI_Obj_Property{
			Name: n,
			Tmln: t,
			User: u,
		}

		cur, err = json.Marshal(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var pat []byte
	{
		var p []map[string]string

		for _, j := range req.Obj.Jsnpatch {
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

	var des []byte
	{
		patch, err := jsonpatch.DecodePatch(pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		des, err = patch.Apply(cur)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var obj audience.CreateI_Obj_Property
	{
		err := json.Unmarshal(des, &obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd bool
	{
		k := fmt.Sprintf(key.Audience, oid)
		e := element.Join(aid, obj.Name, obj.Tmln, obj.User)
		s := aid
		i := index.New(index.Name, obj.Name)

		upd, err = u.redigo.Sorted().Update().Value(k, e, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.UpdateO
	{
		res = &audience.UpdateO{
			Obj: &audience.UpdateO_Obj{
				Metadata: map[string]string{},
			},
		}

		if upd {
			res.Obj.Metadata[metadata.AudienceStatus] = "updated"
		}
	}

	return res, nil
}
