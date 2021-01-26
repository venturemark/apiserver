package creator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/message/element"
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

	var usr string
	{
		usr = req.Obj.Metadata[metadata.UserID]
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

	// We store messages in a sorted set. The elements of the sorted set are
	// concatenated strings of t, m and r. Here t is the unix timestamp
	// referring to the time right now at creation time. Here m is the message
	// text. Here r is the reply ID, if any. We track t as part of the element
	// within the sorted set to guarantee a unique element.
	{
		k := fmt.Sprintf(key.Message, oid, tid, uid)
		e := element.Join(mid, oid, req.Obj.Property.Text, req.Obj.Property.Reid, usr)
		s := mid

		err = c.redigo.Sorted().Create().Element(k, e, s)
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
