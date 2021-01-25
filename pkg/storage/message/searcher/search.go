package searcher

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/message/element"
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
		k := fmt.Sprintf(key.Message, oid, tid, uid)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.SearchO
	{
		res = &message.SearchO{}

		for _, s := range str {
			mid, tex, rid, err := element.Split(s)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &message.SearchO_Obj{
				Metadata: map[string]string{
					metadata.MessageID:      strconv.Itoa(int(mid)),
					metadata.OrganizationID: oid,
					metadata.TimelineID:     tid,
					metadata.UpdateID:       uid,
				},
				Property: &message.SearchO_Obj_Property{
					Text: tex,
					Reid: rid,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
