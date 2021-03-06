package creator

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

// Create provides a storage primitive to persist messages associated with an
// update.
func (c *Creator) Create(req *message.CreateI) (*message.CreateO, error) {
	var err error

	var mei float64
	{
		mei, err = strconv.ParseFloat(req.Obj.Metadata[metadata.MessageID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tii string
	{
		tii = req.Obj.Metadata[metadata.TimelineID]
	}

	var upi string
	{
		upi = req.Obj.Metadata[metadata.UpdateID]
	}

	var vei string
	{
		vei = req.Obj.Metadata[metadata.VentureID]
	}

	var val string
	{
		mes := schema.Message{
			Obj: schema.MessageObj{
				Metadata: req.Obj.Metadata,
				Property: schema.MessageObjProperty{
					Text: req.Obj.Property.Text,
					Reid: req.Obj.Property.Reid,
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
		k := fmt.Sprintf(key.Message, vei, tii, upi)
		v := val
		s := mei

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.CreateO
	{
		res = &message.CreateO{
			Obj: &message.CreateO_Obj{
				Metadata: map[string]string{
					metadata.MessageID: req.Obj.Metadata[metadata.MessageID],
				},
			},
		}
	}

	return res, nil
}
