package creator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/index"
	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/schema"
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
		k := fmt.Sprintf(key.Audience, oid)
		v := val
		s := aid
		i := index.New(index.Name, req.Obj.Property.Name)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
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
