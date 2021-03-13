package searcher

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *venture.SearchI) (*venture.SearchO, error) {
	var err error

	var vek *key.Key
	{
		vek = key.Venture(req.Obj[0].Metadata)
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str string
	{
		k := vek.Elem()

		str, err = s.redigo.Simple().Search().Value(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *venture.SearchO
	{
		res = &venture.SearchO{}

		ven := &schema.Venture{}
		err := json.Unmarshal([]byte(str), ven)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		o := &venture.SearchO_Obj{
			Metadata: ven.Obj.Metadata,
			Property: &venture.SearchO_Obj_Property{
				Desc: ven.Obj.Property.Desc,
				Link: lin(ven.Obj.Property.Link),
				Name: ven.Obj.Property.Name,
			},
		}

		res.Obj = append(res.Obj, o)
	}

	return res, nil
}

func lin(i []schema.VentureObjPropertyLink) []*venture.SearchO_Obj_Property_Link {
	var o []*venture.SearchO_Obj_Property_Link

	for _, l := range i {
		o = append(o, &venture.SearchO_Obj_Property_Link{Addr: l.Addr, Text: l.Text})
	}

	return o
}
