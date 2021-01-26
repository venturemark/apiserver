package creator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	mel "github.com/venturemark/apiserver/pkg/value/metric/element"
	uel "github.com/venturemark/apiserver/pkg/value/update/element"
)

// Create provides a storage primitive to persist metric updates associated with
// a timeline. A timeline refers to many metrics. A timeline does also refer to
// many updates. Between metrics and updates there is a one to one relationship.
// Metrics and updates can be found considering their metadata and time of
// creation. For more information about technical details see the inline
// documentation.
func (c *Creator) Create(req *metupd.CreateI) (*metupd.CreateO, error) {
	var err error

	var aid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// nano seconds, normalized to the UTC timezone. Persisting metrics and
	// updates respectively uses the same timestamp. This is then how we
	// associate one with the other. This is then also how our discovery
	// mechanisms are designed. Everything starts with time, which means that
	// pseudo random IDs are irrelevant for us. Note that we tracked IDs once in
	// seconds, which caused problems when progammatically faking demo
	// timelines, because only one timeline per second could be created.
	var mid float64
	{
		mid = float64(time.Now().UTC().UnixNano())
	}

	// We store metrics in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is a composision of any
	// datapoint relevant to the user on the associated dimensional space. We
	// track t as part of the element within the sorted set to guarantee a
	// unique element, even if the user's coordinates on a timeline ever appear
	// twice.
	{
		k := fmt.Sprintf(key.Metric, aid, tid)
		e := mel.Join(mid, toInterface(req.Obj.Property.Data))
		s := mid

		err = c.redigo.Sorted().Create().Element(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store updates in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is the user's natural
	// language in written form. We track t as part of the element within the
	// sorted set to guarantee a unique element, even if the user's coordinates
	// on a timeline ever appear twice.
	{
		k := fmt.Sprintf(key.Update, aid, tid)
		e := uel.Join(mid, "", req.Obj.Property.Text, "")
		s := mid

		err = c.redigo.Sorted().Create().Element(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *metupd.CreateO
	{
		res = &metupd.CreateO{
			Obj: &metupd.CreateO_Obj{
				Metadata: map[string]string{
					metadata.MetricID: strconv.Itoa(int(mid)),
					metadata.UpdateID: strconv.Itoa(int(mid)),
				},
			},
		}
	}

	return res, nil
}

func toInterface(dat []*metupd.CreateI_Obj_Property_Data) []mel.Interface {
	var l []mel.Interface

	for _, d := range dat {
		l = append(l, d)
	}

	return l
}
