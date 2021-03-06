package creator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

// Create provides a storage primitive to persist timelines associated with a
// user.
func (c *Creator) Create(req *timeline.CreateI) (*timeline.CreateO, error) {
	var err error

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// nano seconds, normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on this very unix timestamp. Everything starts with time,
	// which means that pseudo random IDs are irrelevant for us. Note that we
	// tracked IDs once in seconds, which caused problems when progammatically
	// faking demo timelines, because only one timeline per second could be
	// created.
	var tii float64
	{
		tii = float64(time.Now().UTC().UnixNano())
	}

	var vei string
	{
		vei = req.Obj.Metadata[metadata.VentureID]
	}

	{
		req.Obj.Metadata[metadata.TimelineID] = strconv.FormatFloat(tii, 'f', -1, 64)
	}

	var val string
	{
		tim := schema.Timeline{
			Obj: schema.TimelineObj{
				Metadata: req.Obj.Metadata,
				Property: schema.TimelineObjProperty{
					Desc: req.Obj.Property.Desc,
					Name: req.Obj.Property.Name,
					Stat: "active",
				},
			},
		}

		byt, err := json.Marshal(tim)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := fmt.Sprintf(key.Timeline, vei)
		v := val
		s := tii
		i := index.New(index.Name, req.Obj.Property.Name)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.CreateO
	{
		res = &timeline.CreateO{
			Obj: &timeline.CreateO_Obj{
				Metadata: map[string]string{
					metadata.TimelineID: req.Obj.Metadata[metadata.TimelineID],
				},
			},
		}
	}

	return res, nil
}
