package creator

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *invite.CreateI) (*invite.CreateO, error) {
	var err error

	var ink *key.Key
	{
		ink = key.Invite(req.Obj[0].Metadata)
	}

	var inc string
	{
		inc = req.Obj[0].Metadata[metadata.InviteCode]
	}

	var val string
	{
		tim := schema.Invite{
			Obj: schema.InviteObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.InviteObjProperty{
					Mail: req.Obj[0].Property.Mail,
					Stat: "pending",
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
		k := ink.List()
		v := val
		s := ink.ID().F()
		i := index.New(index.Mail, req.Obj[0].Property.Mail)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *invite.CreateO
	{
		res = &invite.CreateO{
			Obj: []*invite.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.InviteCode: inc,
						metadata.InviteID:   ink.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}
