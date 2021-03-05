package creator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"
)

// Create provides a storage primitive to persist text updates associated with a
// timeline. A timeline refers to many updates. Updates can be found considering
// their metadata and time of creation. For more information about technical
// details see the inline documentation.
func (c *Creator) Create(req *texupd.CreateI) (*texupd.CreateO, error) {
	var err error

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// nano seconds, normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on the unix timestamp, which acts ad ID. Everything starts
	// with time, which means that pseudo random IDs are irrelevant for us. Note
	// that we tracked IDs once in seconds, which caused problems when
	// progammatically faking demo timelines, because only one timeline per
	// second could be created.
	var uid float64
	{
		uid = float64(time.Now().UTC().UnixNano())
	}

	var vid string
	{
		vid = req.Obj.Metadata[metadata.VentureID]
	}

	{
		req.Obj.Metadata[metadata.UpdateID] = strconv.FormatFloat(uid, 'f', -1, 64)
	}

	var val string
	{
		upd := schema.Update{
			Obj: schema.UpdateObj{
				Metadata: req.Obj.Metadata,
				Property: schema.UpdateObjProperty{
					Text: req.Obj.Property.Text,
				},
			},
		}

		byt, err := json.Marshal(upd)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := fmt.Sprintf(key.Update, vid, tid)
		v := val
		s := uid

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.CreateO
	{
		res = &texupd.CreateO{
			Obj: &texupd.CreateO_Obj{
				Metadata: map[string]string{
					metadata.UpdateID: req.Obj.Metadata[metadata.UpdateID],
				},
			},
		}
	}

	return res, nil
}
