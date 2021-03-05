package creator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

// Create provides a storage primitive to persist messages associated with an
// update.
func (c *Creator) Create(req *message.CreateI) (*message.CreateO, error) {
	var err error

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// nano seconds, normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on this very unix timestamp. Everything starts with time,
	// which means that pseudo random IDs are irrelevant for us. Note that we
	// tracked IDs once in seconds, which caused problems when progammatically
	// faking demo timelines, because only one timeline per second could be
	// created.
	var mid float64
	{
		mid = float64(time.Now().UTC().UnixNano())
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	var uid string
	{
		uid = req.Obj.Metadata[metadata.UpdateID]
	}

	var vid string
	{
		vid = req.Obj.Metadata[metadata.VentureID]
	}

	{
		req.Obj.Metadata[metadata.MessageID] = strconv.FormatFloat(mid, 'f', -1, 64)
	}

	var val string
	{
		mes := schema.Message{
			Obj: schema.MessageObj{
				Metadata: req.Obj.Metadata,
				Property: schema.MessageObjProperty{
					Text: req.Obj.Property.Text,
					Reid: req.Obj.Property.Reid,
				},
			},
		}

		byt, err := json.Marshal(mes)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := fmt.Sprintf(key.Message, vid, tid, uid)
		v := val
		s := mid

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.CreateO
	{
		res = &message.CreateO{
			Obj: &message.CreateO_Obj{
				Metadata: map[string]string{
					metadata.MessageID: req.Obj.Metadata[metadata.MessageID],
				},
			},
		}
	}

	return res, nil
}
