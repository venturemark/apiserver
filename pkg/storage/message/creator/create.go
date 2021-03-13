package creator

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *message.CreateI) (*message.CreateO, error) {
	var err error

	var mek *key.Key
	{
		mek = key.Message(req.Obj[0].Metadata)
	}

	var val string
	{
		mes := schema.Message{
			Obj: schema.MessageObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.MessageObjProperty{
					Text: req.Obj[0].Property.Text,
					Reid: req.Obj[0].Property.Reid,
				},
			},
		}

		byt, err := json.Marshal(mes)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := mek.List()
		v := val
		s := mek.ID().F()

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.CreateO
	{
		res = &message.CreateO{
			Obj: []*message.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.MessageID: mek.ID().S(),
					},
				},
			},
		}
	}

	return res, nil
}
