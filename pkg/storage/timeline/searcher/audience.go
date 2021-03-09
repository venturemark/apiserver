package searcher

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) searchAud(req *timeline.SearchI) (*audience.SearchO, error) {
	var err error

	var auk *key.Key
	{
		auk = key.Audience(req.Obj[0].Metadata)
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := auk.List()
		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.SearchO
	{
		res = &audience.SearchO{}

		for _, s := range str {
			aud := &schema.Audience{}
			err := json.Unmarshal([]byte(s), aud)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &audience.SearchO_Obj{
				Metadata: aud.Obj.Metadata,
				Property: &audience.SearchO_Obj_Property{
					Name: aud.Obj.Property.Name,
					Tmln: aud.Obj.Property.Tmln,
					User: aud.Obj.Property.User,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
