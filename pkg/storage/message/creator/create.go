package creator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/schema"
)

// Create provides a storage primitive to persist messages associated with an
// update.
func (c *Creator) Create(req *message.CreateI) (*message.CreateO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	var uid string
	{
		uid = req.Obj.Metadata[metadata.UpdateID]
	}

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
		k := fmt.Sprintf(key.Message, oid, tid, uid)
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
					metadata.MessageID: strconv.Itoa(int(mid)),
				},
			},
		}
	}

	return res, nil
}
