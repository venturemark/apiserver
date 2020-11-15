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
	// while having support for the "bet" operator later. One example is to show
	// metrics within a certain timerange.
	//
	// The data structure of the sorted set looks schematically similar to the
	// example below.
	//
	//     tml:tml-al9qy:met    [n,y,y] [n,y,y] ...
	//
	var str []string
	{
		k := fmt.Sprintf("tml:%s:met", obj.Filter.Property[0].Timeline)

		str, err = t.redigo.Scored().Search(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store metrics in a sorted set. The elements of the sorted set are
	// concatenated strings of n and potentially multiple y coordinates. Here n
	// is the unix timestamp of metric creation. Here any y coordinate
	// represents a datapoint relevant to the user.
	var res []*metric.SearchO_Result
	for _, s := range str {
		now, yaxis, err := splitElement(s)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		r := &metric.SearchO_Result{
			Yaxis:     yaxis,
			Timeline:  obj.Filter.Property[0].Timeline,
			Timestamp: now,
		}

		res = append(res, r)
	}

	return &metric.SearchO{Result: res}, nil
}

func splitElement(s string) (int64, []int64, error) {
	l := strings.Split(s, ",")

	var n int64
	{
		i, err := strconv.Atoi(l[0])
		if err != nil {
			return 0, nil, tracer.Mask(err)
		}

		n = int64(i)
	}

	var y []int64
	for _, p := range l[1:] {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 0, nil, tracer.Mask(err)
		}

		y = append(y, int64(i))
	}

	return n, y, nil
}
