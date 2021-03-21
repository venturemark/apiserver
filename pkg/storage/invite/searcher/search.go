package searcher

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *invite.SearchI) (*invite.SearchO, error) {
	var err error

	var ink *key.Key
	{
		ink = key.Invite(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := ink.List()

		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *invite.SearchO
	{
		res = &invite.SearchO{}

		for _, s := range str {
			inv := &schema.Invite{}
			err := json.Unmarshal([]byte(s), inv)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &invite.SearchO_Obj{
				Metadata: inv.Obj.Metadata,
				Property: &invite.SearchO_Obj_Property{
					Mail: inv.Obj.Property.Mail,
					Stat: inv.Obj.Property.Stat,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
