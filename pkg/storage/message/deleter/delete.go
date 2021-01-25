package deleter

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
)

// Delete provides a storage primitive to remove messages associated with an
// update.
func (c *Deleter) Delete(req *message.DeleteI) (*message.DeleteO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var mid float64
	{
		mid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.MessageID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	var uid string
	{
		uid = req.Obj.Metadata[metadata.UpdateID]
	}

	{
		k := fmt.Sprintf(key.Message, oid, tid, uid)
		s := mid

		err = c.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.DeleteO
	{
		res = &message.DeleteO{
			Obj: &message.DeleteO_Obj{
				Metadata: map[string]string{
					metadata.MessageStatus: "deleted",
				},
			},
		}
	}

	return res, nil
}
