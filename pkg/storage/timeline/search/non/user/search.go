package user

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/timeline/element"
)

// Search provides a filter primitive to lookup timelines associated with a
// user.
func (t *User) Search(req *timeline.SearchI) (*timeline.SearchO, error) {
	var err error

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := fmt.Sprintf(key.UserTimeline, req.Obj[0].Metadata[metadata.User])
		str, err = t.redigo.Scored().Search(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of the unix timestamp of timeline creation and the
	// timeline name.
	var res *timeline.SearchO
	{
		res = &timeline.SearchO{}

		for _, s := range str {
			uni, val, err := element.Split(s)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &timeline.SearchO_Obj{
				Metadata: map[string]string{
					metadata.Unixtime: strconv.Itoa(int(uni)),
					metadata.User:     req.Obj[0].Metadata[metadata.User],
				},
				Property: &timeline.SearchO_Obj_Property{
					Name: val,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
