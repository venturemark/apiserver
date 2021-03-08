package deleter

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

// Delete provides a storage primitive to remove messages associated with an
// update.
func (c *Deleter) Delete(req *message.DeleteI) (*message.DeleteO, error) {
	var err error

	var mei float64
	{
		mei, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.MessageID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tii string
	{
		tii = req.Obj[0].Metadata[metadata.TimelineID]
	}

	var upi string
	{
		upi = req.Obj[0].Metadata[metadata.UpdateID]
	}

	var vei string
	{
		vei = req.Obj[0].Metadata[metadata.VentureID]
	}

	{
		k := fmt.Sprintf(key.Message, vei, tii, upi)
		s := mei

		err = c.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.DeleteO
	{
		res = &message.DeleteO{
			Obj: []*message.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.MessageStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
