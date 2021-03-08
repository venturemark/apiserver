package creator

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	var tii float64
	{
		tii, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.TimelineID], 64)
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
		tim := schema.Timeline{
			Obj: schema.TimelineObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.TimelineObjProperty{
					Desc: req.Obj[0].Property.Desc,
					Name: req.Obj[0].Property.Name,
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
		i := index.New(index.Name, req.Obj[0].Property.Name)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.CreateO
	{
		res = &timeline.CreateO{
			Obj: []*timeline.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.TimelineID: req.Obj[0].Metadata[metadata.TimelineID],
					},
				},
			},
		}
	}

	return res, nil
}
