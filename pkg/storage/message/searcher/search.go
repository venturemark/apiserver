package searcher

import (
	"encoding/json"
	"fmt"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

// Search provides a filter primitive to lookup messages associated with an
// update.
func (s *Searcher) Search(req *message.SearchI) (*message.SearchO, error) {
	var err error

	var oid string
	{
		oid = req.Obj[0].Metadata[metadata.OrganizationID]
	}

	var tid string
	{
		tid = req.Obj[0].Metadata[metadata.TimelineID]
	}

	var uid string
	{
		uid = req.Obj[0].Metadata[metadata.UpdateID]
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := fmt.Sprintf(key.MessageResource, oid, tid, uid)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
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
