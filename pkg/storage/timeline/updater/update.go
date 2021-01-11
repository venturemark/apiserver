package updater

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/index"
	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/timeline/element"
)

// Update provides a storage primitive to modify timelines associated with a
// user.
func (u *Updater) Update(req *timeline.UpdateI) (*timeline.UpdateO, error) {
	var err error

	var aid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]
	}

	var tid float64
	{
		tid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.TimelineID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is the timeline name. We
	// track t as part of the element within the sorted set to guarantee a
	// unique element.
	var upd bool
	{
		k := fmt.Sprintf(key.Timeline, aid)
		e := element.Join(tid, req.Obj.Property.Name)
		s := tid
		i := index.New(index.Name, req.Obj.Property.Name)

		upd, err = u.redigo.Sorted().Update().Value(k, e, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.UpdateO
	{
		res = &timeline.UpdateO{
			Obj: &timeline.UpdateO_Obj{
				Metadata: map[string]string{},
			},
		}

		if upd {
			res.Obj.Metadata[metadata.TimelineStatus] = "updated"
		}
	}

	return res, nil
}
