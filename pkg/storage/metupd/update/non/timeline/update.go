package timeline

import (
	"fmt"
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"
)

// Update provides a storage primitive to modify metric updates associated with
// a timeline. A timeline refers to many metrics. A timeline does also refer to
// many updates. Between metrics and updates there is a one to one relationship.
// Metrics and updates can be found considering their scope and time of
// creation. For more information about technical details see the inline
// documentation.
func (t *Timeline) Update(obj *metupd.UpdateI) (*metupd.UpdateO, error) {
	var err error

	// When updating metric updates all assumptions are equal to creating metric
	// updates. The update mechanism of elements within sorted sets is rather
	// complex. An error will be returned if the sorted set or its alleged
	// element does not exist. The bool met will be false if the metrics to
	// update did in fact not change. Then no update will be performed under the
	// hood. Note that we should only perform the update if we are certain there
	// was information provided for the update in the first place.
	var met bool
	if len(obj.Yaxis) != 0 {
		k := fmt.Sprintf("tml:%s:met", obj.Timeline)
		e := joinYaxis(obj.Timestamp, obj.Yaxis...)
		s := float64(obj.Timestamp)

		met, err = t.redigo.Scored().Update(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// When updating metric updates all assumptions are equal to creating metric
	// updates. The update mechanism of elements within sorted sets is rather
	// complex. An error will be returned if the sorted set or its alleged
	// element does not exist. The bool upd will be false if the update to
	// update did in fact not change. Then no update will be performed under the
	// hood. Note that we should only perform the update if we are certain there
	// was information provided for the update in the first place.
	var upd bool
	if obj.Text != "" {
		k := fmt.Sprintf("tml:%s:upd", obj.Timeline)
		e := fmt.Sprintf("%d,%s", obj.Timestamp, obj.Text)
		s := float64(obj.Timestamp)

		upd, err = t.redigo.Scored().Update(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return &metupd.UpdateO{Metric: met, Update: upd}, nil
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
