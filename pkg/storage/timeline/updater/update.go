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

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var tid float64
	{
		tid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.TimelineID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Properties can be updated separately. We only need to figure out which
	// information are given so that we only update the specified fields. For
	// any field not being specified within the update request we fetch the
	// current state and default to it accordingly.
	var des string
	var nam string
	var sta string
	{
		k := fmt.Sprintf(key.Timeline, oid)

		str, err := u.redigo.Sorted().Search().Score(k, tid, tid)
		if err != nil {
			return nil, tracer.Mask(err)
		}
		_, d, n, s, err := element.Split(str[0])
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if req.Obj.Property.Desc != nil {
			des = *req.Obj.Property.Desc
		} else {
			des = d
		}

		if req.Obj.Property.Name != nil {
			nam = *req.Obj.Property.Name
		} else {
			nam = n
		}

		if req.Obj.Property.Stat != nil {
			sta = *req.Obj.Property.Stat
		} else {
			sta = s
		}
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is the timeline name. We
	// track t as part of the element within the sorted set to guarantee a
	// unique element.
	var upd bool
	{
		k := fmt.Sprintf(key.Timeline, oid)
		e := element.Join(tid, des, nam, sta)
		s := tid

		var i string
		if nam != "" {
			// If the name got updated we need to update its index as well.
			i = index.New(index.Name, nam)
		}

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
