package creator

import (
	"encoding/json"

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

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
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
		k := tik.List()
		v := val
		s := tik.ID().F()
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
						metadata.TimelineID: tik.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}
