package updater

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	mel "github.com/venturemark/apiserver/pkg/value/metric/element"
	uel "github.com/venturemark/apiserver/pkg/value/update/element"
)

// Update provides a storage primitive to modify metric updates associated with
// a timeline. A timeline refers to many metrics. A timeline does also refer to
// many updates. Between metrics and updates there is a one to one relationship.
// Metrics and updates can be found considering their scope and time of
// creation. For more information about technical details see the inline
// documentation.
func (u *Updater) Update(req *metupd.UpdateI) (*metupd.UpdateO, error) {
	var aid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
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
		i, err := strconv.ParseFloat(req.Obj.Metadata[metadata.MetricID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		k := fmt.Sprintf(key.Metric, aid, tid)
		e := mel.Join(i, toInterface(req.Obj.Property.Data))
		s := i

		met, err = u.redigo.Sorted().Update().Value(k, e, s)
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
		i, err := strconv.ParseFloat(req.Obj.Metadata[metadata.UpdateID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		k := fmt.Sprintf(key.Update, aid, tid)
		e := uel.Join(i, req.Obj.Property.Text)
		s := i

		upd, err = u.redigo.Sorted().Update().Value(k, e, s)
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

func toInterface(dat []*metupd.UpdateI_Obj_Property_Data) []mel.Interface {
	var l []mel.Interface

	for _, d := range dat {
		l = append(l, d)
	}

	return l
}
