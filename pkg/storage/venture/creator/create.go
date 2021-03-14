package creator

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *venture.CreateI) (*venture.CreateO, error) {
	var err error

	var vek *key.Key
	{
		vek = key.Venture(req.Obj[0].Metadata)
	}

	var val string
	{
		ven := schema.Venture{
			Obj: schema.VentureObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.VentureObjProperty{
					Desc: req.Obj[0].Property.Desc,
					Link: lin(req.Obj[0].Property.Link),
					Name: req.Obj[0].Property.Name,
				},
			},
		}

		byt, err := json.Marshal(ven)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := vek.Elem()
		v := val

		err = c.redigo.Simple().Create().Element(k, v)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *venture.CreateO
	{
		res = &venture.CreateO{
			Obj: []*venture.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.VentureID: vek.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}

func lin(i []*venture.CreateI_Obj_Property_Link) []schema.VentureObjPropertyLink {
	var o []schema.VentureObjPropertyLink

	for _, l := range i {
		o = append(o, schema.VentureObjPropertyLink{Addr: l.Addr, Text: l.Text})
	}

	return o
}
