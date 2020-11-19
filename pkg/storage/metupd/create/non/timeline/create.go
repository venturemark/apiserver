package timeline

import (
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/metric/timeline/data"
)

// Create provides a storage primitive to persist metric updates associated with
// a timeline. A timeline refers to many metrics. A timeline does also refer to
// many updates. Between metrics and updates there is a one to one relationship.
// Metrics and updates can be found considering their metadata and time of
// creation. For more information about technical details see the inline
// documentation.
func (t *Timeline) Create(req *metupd.CreateI) (*metupd.CreateO, error) {
	var err error

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// seconds is normalized to the UTC timezone. Persisting metrics and updates
	// respectively uses the same timestamp. This is then how we associate one
	// with the other. This is then also how our discovery mechanisms are
	// designed. Everything starts with time, which means that pseudo random IDs
	// are irrelevant for us.
	var now int64
	{
		now = time.Now().UTC().Unix()
	}

	// We store metrics in a sorted set. The elements of the sorted set are
	// concatenated strings of n and potentially multiple datapoints per
	// dimensional space. Here n is the unix timestamp referring to the time
	// right now at creation time. Here any y value represents a datapoint
	// relevant to the user on the associated dimensional space. We track n as
	// part of the element within the sorted set to guarantee a unique element,
	// even if the user's coordinates on a timeline ever appear twice.
	{
		k := fmt.Sprintf(key.TimelineMetric, req.Obj.Metadata[metadata.Timeline])
		e := data.Join(now, toInterface(req.Obj.Property.Data))
		s := float64(now)

		err = t.redigo.Scored().Create(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store updates in a sorted set. The elements of the sorted set are
	// concatenated strings of n and t. Here n is the unix timestamp referring
	// to the time right now at creation time. Here t is the user's natural
	// language in written form. We track n as part of the element within the
	// sorted set to guarantee a unique element, even if the user's coordinates
	// on a timeline ever appear twice.
	{
		k := fmt.Sprintf(key.TimelineUpdate, req.Obj.Metadata[metadata.Timeline])
		e := fmt.Sprintf("%d,%s", now, req.Obj.Property.Text)
		s := float64(now)

		err = t.redigo.Scored().Create(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *metupd.CreateO
	{
		res = &metupd.CreateO{
			Obj: &metupd.CreateO_Obj{
				Metadata: map[string]string{
					metadata.Unixtime: strconv.Itoa(int(now)),
				},
			},
		}
	}

	return res, nil
}

func toInterface(dat []*metupd.CreateI_Obj_Property_Data) []data.Interface {
	var l []data.Interface

	for _, d := range dat {
		l = append(l, d)
	}

	return l
}
