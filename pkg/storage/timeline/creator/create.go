package creator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/timeline/element"
)

// Create provides a storage primitive to persist timelines associated with a
// user.
func (c *Creator) Create(req *timeline.CreateI) (*timeline.CreateO, error) {
	var err error

	var usr string
	{
		usr = req.Obj.Metadata[metadata.UserID]
	}

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// seconds is normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on this very unix timestamp. Everything starts with time,
	// which means that pseudo random IDs are irrelevant for us.
	var uni float64
	{
		uni = float64(time.Now().UTC().Unix())
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of t and e. Here t is the unix timestamp referring
	// to the time right now at creation time. Here e is the timeline name. We
	// track t as part of the element within the sorted set to guarantee a
	// unique element.
	{
		k := fmt.Sprintf(key.Timeline, usr)
		e := element.Join(uni, req.Obj.Property.Name)
		s := uni

		err = c.redigo.Scored().Create(k, e, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.CreateO
	{
		res = &timeline.CreateO{
			Obj: &timeline.CreateO_Obj{
				Metadata: map[string]string{
					metadata.TimelineID: strconv.Itoa(int(uni)),
				},
			},
		}
	}

	return res, nil
}
