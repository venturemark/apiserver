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

// Search provides a filter primitive to lookup timelines associated with a
// user.
func (s *Searcher) Search(req *audience.SearchI) (*audience.SearchO, error) {
	var err error

	var uid string
	{
		uid = req.Obj[0].Metadata[metadata.UserID]
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := fmt.Sprintf(key.Audience, uid)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of the unix timestamp of timeline creation and the
	// timeline name.
	var res *audience.SearchO
	{
		res = &audience.SearchO{}

		for _, s := range str {
			aid, nam, usr, err := element.Split(s)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &audience.SearchO_Obj{
				Metadata: map[string]string{
					metadata.AudienceID: strconv.Itoa(int(aid)),
					metadata.UserID:     uid,
				},
				Property: &audience.SearchO_Obj_Property{
					Name: nam,
					User: usr,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
