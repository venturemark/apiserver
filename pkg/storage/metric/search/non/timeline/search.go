package timeline

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

// Search provides a filter primitive to lookup metrics associated with a
// timeline. A timeline refers to many metrics. Metrics can be found considering
// their scope and time of creation. For more information about technical
// details see the inline documentation.
func (t *Timeline) Search(obj *metric.SearchI) (*metric.SearchO, error) {
	var err error

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	//
	// With redis we use ZRANGEBYSCORE which allows us to search for objects
	// while having support for the "bet" operator. One example is to show
	// metrics within a certain timerange.
	//
	//
	//     tml:tml-al9qy:met    [n,x,y] [n,x,y] ...
	//
	var str []string
	{
		k := fmt.Sprintf("tml:%s:met", obj.Filter.Property[0].Timeline)

		str, err = t.redigo.Scored().Search(k, 1000)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store metrics in a sorted set. The elements of the sorted set are
	// concatenated strings of n, x and y. Here n is the unix timestamp. Here x
	// by convention is the datapoint of the x axis of a graph. Here y by
	// convention is the datapoint of the y axis of a graph. The scores of the
	// sorted set are unix timestamps.
	var res []*metric.SearchO_Result
	for i := 0; i < len(str); i += 2 {
		l := strings.Split(str[i], ",")

		n, err := strconv.Atoi(l[0])
		if err != nil {
			return nil, tracer.Mask(err)
		}
		x, err := strconv.Atoi(l[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}
		y, err := strconv.Atoi(l[2])
		if err != nil {
			return nil, tracer.Mask(err)
		}

		r := &metric.SearchO_Result{
			Datapoint: []int64{
				int64(x),
				int64(y),
			},
			Timeline:  obj.Filter.Property[0].Timeline,
			Timestamp: int64(n),
		}

		res = append(res, r)
	}

	return &metric.SearchO{Result: res}, nil
}
