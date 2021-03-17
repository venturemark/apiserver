package searcher

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *timeline.SearchI) (*timeline.SearchO, error) {
	var err error

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := tik.List()

		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.SearchO
	{
		res = &timeline.SearchO{}

		for _, s := range str {
			tim := &schema.Timeline{}
			err := json.Unmarshal([]byte(s), tim)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &timeline.SearchO_Obj{
				Metadata: tim.Obj.Metadata,
				Property: &timeline.SearchO_Obj_Property{
					Desc: tim.Obj.Property.Desc,
					Name: tim.Obj.Property.Name,
					Stat: tim.Obj.Property.Stat,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
