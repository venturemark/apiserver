package creator

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *audience.CreateI) (*audience.CreateO, error) {
	var err error

	var auk *key.Key
	{
		auk = key.Audience(req.Obj[0].Metadata)
	}

	var val string
	{
		aud := schema.Audience{
			Obj: schema.AudienceObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.AudienceObjProperty{
					Name: req.Obj[0].Property.Name,
					Tmln: req.Obj[0].Property.Tmln,
					User: req.Obj[0].Property.User,
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
		k := auk.List()
		v := val
		s := auk.ID().F()
		i := index.New(index.Name, req.Obj[0].Property.Name)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.CreateO
	{
		res = &audience.CreateO{
			Obj: []*audience.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.AudienceID: auk.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}
