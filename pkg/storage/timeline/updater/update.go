package updater

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
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

	var tim *schema.Timeline
	{
		k := fmt.Sprintf(key.TimelineResource, oid)

		s, err := u.redigo.Sorted().Search().Score(k, tid, tid)
		if err != nil {
			return nil, tracer.Mask(err)
		}
		tim = &schema.Timeline{}
		err = json.Unmarshal([]byte(s[0]), tim)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if req.Obj.Property.Desc != nil {
			tim.Obj.Property.Desc = *req.Obj.Property.Desc
		}

		if req.Obj.Property.Name != nil {
			tim.Obj.Property.Name = *req.Obj.Property.Name
		}

		if req.Obj.Property.Stat != nil {
			tim.Obj.Property.Stat = *req.Obj.Property.Stat
		}
	}

	var val string
	{
		byt, err := json.Marshal(tim)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is the timeline name. We
	// track t as part of the element within the sorted set to guarantee a
	// unique element.
	var upd bool
	{
		k := fmt.Sprintf(key.TimelineResource, oid)
		v := val
		s := tid

		var i string
		if req.Obj.Property.Name != nil {
			// If the name got updated we need to update its index as well.
			i = index.New(index.Name, tim.Obj.Property.Name)
		}

		upd, err = u.redigo.Sorted().Update().Value(k, v, s, i)
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
