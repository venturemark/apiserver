package creator

import (
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/index"
	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/audience/element"
)

func (c *Creator) Create(req *audience.CreateI) (*audience.CreateO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// nano seconds, normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on this very unix timestamp. Everything starts with time,
	// which means that pseudo random IDs are irrelevant for us. Note that we
	// tracked IDs once in seconds, which caused problems when progammatically
	// faking demo timelines, because only one timeline per second could be
	// created.
	var aid float64
	{
		aid = float64(time.Now().UTC().UnixNano())
	}

	// We store audiences in a sorted set. The elements of the sorted set are
	// concatenated strings of t, n and u. Here t is the unix timestamp
	// referring to the time right now at creation time. Here n is the audience
	// name. Here u is a list of users associated to that audience. We track t
	// as part of the element within the sorted set to guarantee a unique
	// element.
	{
		k := fmt.Sprintf(key.Audience, oid)
		e := element.Join(aid, req.Obj.Property.Name, req.Obj.Property.Tmln, req.Obj.Property.User)
		s := aid
		i := index.New(index.Name, req.Obj.Property.Name)

		err = c.redigo.Sorted().Create().Element(k, e, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.CreateO
	{
		res = &audience.CreateO{
			Obj: &audience.CreateO_Obj{
				Metadata: map[string]string{
					metadata.AudienceID: strconv.Itoa(int(aid)),
				},
			},
		}
	}

	return res, nil
}
