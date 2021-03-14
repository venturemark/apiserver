package searcher

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *message.SearchI) (*message.SearchO, error) {
	var err error

	var mek *key.Key
	{
		mek = key.Message(req.Obj[0].Metadata)
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := mek.List()
		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.SearchO
	{
		res = &message.SearchO{}

		for _, s := range str {
			mes := &schema.Message{}
			err := json.Unmarshal([]byte(s), mes)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &message.SearchO_Obj{
				Metadata: mes.Obj.Metadata,
				Property: &message.SearchO_Obj_Property{
					Text: mes.Obj.Property.Text,
					Reid: mes.Obj.Property.Reid,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
