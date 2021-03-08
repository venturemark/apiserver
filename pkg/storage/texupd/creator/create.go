package creator

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	var tii string
	{
		tii = req.Obj[0].Metadata[metadata.TimelineID]
	}

	var upi float64
	{
		upi, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.UpdateID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var vei string
	{
		vei = req.Obj[0].Metadata[metadata.VentureID]
	}

	var val string
	{
		upd := schema.Update{
			Obj: schema.UpdateObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.UpdateObjProperty{
					Text: req.Obj[0].Property.Text,
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
		k := fmt.Sprintf(key.Update, vei, tii)
		v := val
		s := upi

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.CreateO
	{
		res = &texupd.CreateO{
			Obj: []*texupd.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.UpdateID: req.Obj[0].Metadata[metadata.UpdateID],
					},
				},
			},
		}
	}

	return res, nil
}
