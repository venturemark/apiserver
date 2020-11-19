package timeline

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/update"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/update/timeline/data"
)

// Search provides a filter primitive to lookup updates associated with a
// timeline. A timeline refers to many updates. Updates can be found considering
// their scope and time of creation. For more information about technical
// details see the inline documentation.
func (t *Timeline) Search(req *update.SearchI) (*update.SearchO, error) {
	var err error

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	//
	// With redis we use ZRANGEBYSCORE which allows us to search for objects
	// while having support for the "bet" operator later. One example is to show
	// updates within a certain timerange.
	var str []string
	{
		k := fmt.Sprintf(key.TimelineUpdate, req.Obj[0].Metadata[metadata.Timeline])
		str, err = t.redigo.Scored().Search(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store updates in a sorted set. The elements of the sorted set are
	// concatenated strings of the unix timestamp of update creation and the
	// user's natural language in written form.
	var res *update.SearchO
	{
		res = &update.SearchO{}

		for _, s := range str {
			uni, val, err := data.Split(s)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &update.SearchO_Obj{
				Metadata: map[string]string{
					metadata.Timeline: req.Obj[0].Metadata[metadata.Timeline],
					metadata.Unixtime: strconv.Itoa(int(uni)),
				},
				Property: &update.SearchO_Obj_Property{
					Text: val,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
