package creator

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *role.CreateI) (*role.CreateO, error) {
	var err error

	var rok *key.Key
	{
		rok = key.Role(req.Obj[0].Metadata)
	}

	var suk *key.Key
	{
		suk = key.Subject(req.Obj[0].Metadata)
	}

	var val string
	{
		rol := schema.Role{
			Obj: schema.RoleObj{
				Metadata: req.Obj[0].Metadata,
			},
		}

		byt, err := json.Marshal(rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := rok.List()
		v := val
		s := rok.ID().F()
		i := suk.ID().S()

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := suk.Elem()
		v := rok.List() + ":" + rok.ID().S()
		s := rok.ID().F()

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *role.CreateO
	{
		res = &role.CreateO{
			Obj: []*role.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.RoleID: rok.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}
