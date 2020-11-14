package timeline

import (
	"fmt"
	"strings"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"
)

// Create provides a storage primitive to persist metric updates associated with
// a timeline. A timeline refers to many metrics. A timeline does also refer to
// many updates. Between metrics and updates there is a one to one relationship.
// Metrics and updates can be found considering their scope and time of
// creation. For more information about technical details see the inline
// documentation.
func (t *Timeline) Create(obj *metupd.CreateI) (*metupd.CreateO, error) {
	var err error

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// seconds is normalized to the UTC timezone. Persisting metrics and updates
	// uses the same timestamp. This is then how we associate one with the
	// other. This is then also how our discovery mechanisms are designed.
	// Everything starts with time, making pseudo random IDs irrelevant.
	var now int64
	{
		now = time.Now().UTC().Unix()
	}

	// We store metrics in a sorted set. The elements of the sorted set are
	// concatenated strings of n and potentially multiple y coordinates. Here n
	// is the unix timestamp referring to the time right now at creation time.
	// We track n as part of the element within the sorted set to guarantee a
	// unique element, even if the user's coordinates on a timeline ever appear
	// twice. Future considerations should take redis streams into account for
	// having a more suitable datatype. The scores of the sorted set are unix
	// timestamps.
	//
	//     tml:tml-al9qy:met    [n,y,y] [n,y,y] ...
	//
	{
		k := fmt.Sprintf("tml:%s:met", obj.Timeline)
		e := joinYaxis(now, obj.Yaxis...)
		s := float64(now)

		err = t.redigo.Scored().Create(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store updates in a sorted set. The elements of the sorted set are
	// strings of the user's natural language in written form.
	//
	//     tml:tml-al9qy:upd    [Lorem ipsum ...] [Lorem ipsum ...] ...
	//
	{
		k := fmt.Sprintf("tml:%s:upd", obj.Timeline)
		e := obj.Text
		s := float64(now)

		err = t.redigo.Scored().Create(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return &metupd.CreateO{Timestamp: now}, nil
}

func joinYaxis(now int64, yaxis ...int64) string {
	l := []string{
		fmt.Sprintf("%d", now),
	}

	for _, y := range yaxis {
		l = append(l, fmt.Sprintf("%d", y))
	}

	return strings.Join(l, ",")
}
