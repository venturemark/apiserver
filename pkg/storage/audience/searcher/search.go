package searcher

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/audience/element"
)

func (s *Searcher) Search(req *audience.SearchI) (*audience.SearchO, error) {
	var err error

	var oid string
	{
		oid = req.Obj[0].Metadata[metadata.OrganizationID]
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := fmt.Sprintf(key.Audience, oid)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.SearchO
	{
		res = &audience.SearchO{}

		for _, s := range str {
			aid, nam, tim, usr, err := element.Split(s)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &audience.SearchO_Obj{
				Metadata: map[string]string{
					metadata.AudienceID:     strconv.Itoa(int(aid)),
					metadata.OrganizationID: oid,
				},
				Property: &audience.SearchO_Obj_Property{
					Name: nam,
					Tmln: tim,
					User: usr,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
