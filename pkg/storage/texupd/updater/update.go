package updater

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	uel "github.com/venturemark/apiserver/pkg/value/update/element"
)

// Update provides a storage primitive to modify text updates associated with a
// timeline. A timeline refers to many updates. For more information about
// technical details see the inline documentation.
func (u *Updater) Update(req *texupd.UpdateI) (*texupd.UpdateO, error) {
	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	// When updating text updates all assumptions are equal to creating text
	// updates. The update mechanism of elements within sorted sets is rather
	// complex. An error will be returned if the sorted set or its alleged
	// element does not exist. The bool upd will be false if the update
	// requested to be updated did in fact not change. Then no update will be
	// performed under the hood.
	var upd bool
	{
		i, err := strconv.ParseFloat(req.Obj.Metadata[metadata.UpdateID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		k := fmt.Sprintf(key.Update, oid, tid)
		e := uel.Join(i, req.Obj.Property.Text)
		s := i

		upd, err = u.redigo.Sorted().Update().Value(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.UpdateO
	{
		res = &texupd.UpdateO{
			Obj: &texupd.UpdateO_Obj{
				Metadata: map[string]string{},
			},
		}

		if upd {
			res.Obj.Metadata[metadata.UpdateStatus] = "updated"
		}
	}

	return res, nil
}
