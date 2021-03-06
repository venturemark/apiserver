package creator

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *role.CreateI) (*role.CreateO, error) {
	var err error

	var rei string
	{
		rei = req.Obj[0].Metadata[metadata.ResourceID]
	}

	var roi float64
	{
		roi, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.RoleID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sui string
	{
		sui = req.Obj[0].Metadata[metadata.SubjectID]
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
		k := fmt.Sprintf(key.Role, rei)
		v := val
		s := roi
		i := sui

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
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
						metadata.RoleID: req.Obj[0].Metadata[metadata.RoleID],
					},
				},
			},
		}
	}

	return res, nil
}
