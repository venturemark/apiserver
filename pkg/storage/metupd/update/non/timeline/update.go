package timeline

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	mdat "github.com/venturemark/apiserver/pkg/value/metric/timeline/data"
	udat "github.com/venturemark/apiserver/pkg/value/update/timeline/data"
)

// Update provides a storage primitive to modify metric updates associated with
// a timeline. A timeline refers to many metrics. A timeline does also refer to
// many updates. Between metrics and updates there is a one to one relationship.
// Metrics and updates can be found considering their scope and time of
// creation. For more information about technical details see the inline
// documentation.
func (t *Timeline) Update(req *metupd.UpdateI) (*metupd.UpdateO, error) {
	var err error

	var uni float64
	{
		uni, err = strconv.ParseFloat(req.Obj.Metadata[metadata.Unixtime], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// When updating metric updates all assumptions are equal to creating metric
	// updates. The update mechanism of elements within sorted sets is rather
	// complex. An error will be returned if the sorted set or its alleged
	// element does not exist. The bool met will be false if the metrics to
	// update did in fact not change. Then no update will be performed under the
	// hood. Note that we should only perform the update if we are certain there
	// was information provided for the update in the first place.
	var met bool
	if len(req.Obj.Property.Data) != 0 {
		k := fmt.Sprintf(key.TimelineMetric, req.Obj.Metadata[metadata.Timeline])
		e := mdat.Join(uni, toInterface(req.Obj.Property.Data))
		s := uni

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
	if req.Obj.Property.Text != "" {
		k := fmt.Sprintf(key.TimelineUpdate, req.Obj.Metadata[metadata.Timeline])
		e := udat.Join(uni, req.Obj.Property.Text)
		s := uni

		upd, err = t.redigo.Scored().Update(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *metupd.UpdateO
	{
		res = &metupd.UpdateO{
			Obj: &metupd.UpdateO_Obj{
				Metadata: map[string]string{},
			},
		}

		if met {
			res.Obj.Metadata[metadata.MetricStatus] = "updated"
		}
		if upd {
			res.Obj.Metadata[metadata.UpdateStatus] = "updated"
		}
	}

	return res, nil
}

func toInterface(dat []*metupd.UpdateI_Obj_Property_Data) []mdat.Interface {
	var l []mdat.Interface

	for _, d := range dat {
		l = append(l, d)
	}

	return l
}
