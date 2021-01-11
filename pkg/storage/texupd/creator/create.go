package creator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	uel "github.com/venturemark/apiserver/pkg/value/update/element"
)

// Create provides a storage primitive to persist text updates associated with a
// timeline. A timeline refers to many updates. Updates can be found considering
// their metadata and time of creation. For more information about technical
// details see the inline documentation.
func (c *Creator) Create(req *texupd.CreateI) (*texupd.CreateO, error) {
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
	// nano seconds, normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on the unix timestamp, which acts ad ID. Everything starts
	// with time, which means that pseudo random IDs are irrelevant for us. Note
	// that we tracked IDs once in seconds, which caused problems when
	// progammatically faking demo timelines, because only one timeline per
	// second could be created.
	var tui float64
	{
		tui = float64(time.Now().UTC().UnixNano())
	}

	// We store updates in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is the user's natural
	// language in written form. We track t as part of the element within the
	// sorted set to guarantee a unique element, even if the user's coordinates
	// on a timeline ever appear twice.
	{
		k := fmt.Sprintf(key.Update, aid, tid)
		e := uel.Join(tui, req.Obj.Property.Text)
		s := tui

		err = c.redigo.Sorted().Create().Element(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.CreateO
	{
		res = &texupd.CreateO{
			Obj: &texupd.CreateO_Obj{
				Metadata: map[string]string{
					metadata.UpdateID: strconv.Itoa(int(tui)),
				},
			},
		}
	}

	return res, nil
}
