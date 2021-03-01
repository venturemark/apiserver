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
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *audience.CreateI) (*audience.CreateO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
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
	var aid float64
	{
		aid = float64(time.Now().UTC().UnixNano())
	}

	{
		req.Obj.Metadata[metadata.AudienceID] = strconv.FormatFloat(aid, 'f', -1, 64)
	}

	var val string
	{
		aud := schema.Audience{
			Obj: schema.AudienceObj{
				Metadata: req.Obj.Metadata,
				Property: schema.AudienceObjProperty{
					Name: req.Obj.Property.Name,
					Tmln: req.Obj.Property.Tmln,
					User: req.Obj.Property.User,
				},
			},
		}

		byt, err := json.Marshal(aud)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := fmt.Sprintf(key.AudienceResource, oid)
		v := val
		s := aid
		i := index.New(index.Name, req.Obj.Property.Name)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//     rx    o1 o2
	{
		k := fmt.Sprintf(key.AudienceOwner, oid)
		v := val
		s := aid

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.CreateO
	{
		res = &audience.CreateO{
			Obj: &audience.CreateO_Obj{
				Metadata: map[string]string{
					metadata.AudienceID: req.Obj.Metadata[metadata.AudienceID],
				},
			},
		}
	}

	return res, nil
}
