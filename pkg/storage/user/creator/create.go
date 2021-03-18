package creator

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *user.CreateI) (*user.CreateO, error) {
	var err error

	var clk *key.Key
	{
		clk = key.Claim(req.Obj[0].Metadata)
	}

	var usi string
	{
		usi = req.Obj[0].Metadata[metadata.UserID]
	}

	var usk *key.Key
	{
		usk = key.User(req.Obj[0].Metadata)
	}

	var val string
	{
		ven := schema.User{
			Obj: schema.UserObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.UserObjProperty{
					Desc: req.Obj[0].Property.Desc,
					Prof: lin(req.Obj[0].Property.Prof),
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
		k := usk.Elem()
		v := val

		err = c.redigo.Simple().Create().Element(k, v)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err := c.association.Create(clk, usi)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *user.CreateO
	{
		res = &user.CreateO{
			Obj: []*user.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.UserID: usk.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}

func lin(i []*user.CreateI_Obj_Property_Prof) []schema.UserObjPropertyProf {
	var o []schema.UserObjPropertyProf

	for _, l := range i {
		o = append(o, schema.UserObjPropertyProf{Desc: l.Desc, Vent: l.Vent})
	}

	return o
}
