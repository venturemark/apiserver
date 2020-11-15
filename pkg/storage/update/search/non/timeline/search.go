package timeline

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/update"
	"github.com/xh3b4sd/tracer"
)

// Search provides a filter primitive to lookup updates associated with a
// timeline. A timeline refers to many updates. Updates can be found considering
// their scope and time of creation. For more information about technical
// details see the inline documentation.
func (t *Timeline) Search(obj *update.SearchI) (*update.SearchO, error) {
	var err error

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	//
	// With redis we use ZRANGEBYSCORE which allows us to search for objects
	// while having support for the "bet" operator later. One example is to show
	// updates within a certain timerange.
	//
	// The data structure of the sorted set looks schematically similar to the
	// example below.
	//
	//     tml:tml-al9qy:upd    [n,t] [n,t] ...
	//
	var str []string
	{
		k := fmt.Sprintf("tml:%s:upd", obj.Filter.Property[0].Timeline)

		str, err = t.redigo.Scored().Search(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store updates in a sorted set. The elements of the sorted set are
	// concatenated strings of n and t. Here n is the unix timestamp of update
	// creation. Here t is the user's natural language in written form.
	var res []*update.SearchO_Result
	for _, s := range str {
		now, text, err := splitElement(s)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		r := &update.SearchO_Result{
			Text:      text,
			Timeline:  obj.Filter.Property[0].Timeline,
			Timestamp: now,
		}

		res = append(res, r)
	}

	return &update.SearchO{Result: res}, nil
}

func splitElement(s string) (int64, string, error) {
	l := strings.Split(s, ",")

	var n int64
	{
		i, err := strconv.Atoi(l[0])
		if err != nil {
			return 0, "", tracer.Mask(err)
		}

		n = int64(i)
	}

	var t string
	{
		t = l[1]
	}

	return n, t, nil
}
