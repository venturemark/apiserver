package searcher

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/metric/element"
)

// Search provides a filter primitive to lookup metrics associated with a
// timeline. A timeline refers to many metrics. Metrics can be found considering
// their scope and time of creation. For more information about technical
// details see the inline documentation.
func (s *Searcher) Search(req *metric.SearchI) (*metric.SearchO, error) {
	var err error

	var tml string
	var usr string
	{
		tml = req.Obj[0].Metadata[metadata.TimelineID]
		usr = req.Obj[0].Metadata[metadata.UserID]
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	//
	// With redis we use ZRANGEBYSCORE which allows us to search for objects
	// while having support for the "bet" operator later. One example is to show
	// metrics within a certain timerange.
	var str []string
	{
		k := fmt.Sprintf(key.Metric, usr, tml)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store metrics in a sorted set. The elements of the sorted set are
	// concatenated strings of the unix timestamp of metric creation and
	// potentially multiple datapoints of different dimensional spaces. Note
	// that we include the reserved dimensional space t for the creation time of
	// the datapoints.
	var res *metric.SearchO
	{
		res = &metric.SearchO{}

		for _, s := range str {
			uni, val, err := element.Split(s)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			var dat []*metric.SearchO_Obj_Property_Data
			for _, v := range val {
				d := &metric.SearchO_Obj_Property_Data{
					Space: v.GetSpace(),
					Value: v.GetValue(),
				}

				dat = append(dat, d)
			}

			o := &metric.SearchO_Obj{
				Metadata: map[string]string{
					metadata.MetricID:   strconv.Itoa(int(uni)),
					metadata.TimelineID: tml,
					metadata.UserID:     usr,
				},
				Property: &metric.SearchO_Obj_Property{
					Data: dat,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
