package deleter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

// Delete provides a storage primitive to remove messages associated with an
// update.
func (d *Deleter) Delete(req *message.DeleteI) (*message.DeleteO, error) {
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

	var mes *schema.Message
	{
		k := fmt.Sprintf(key.Message, oid, tid, uid)
		s, err := d.redigo.Sorted().Search().Score(k, mid, mid)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		mes = &schema.Message{}
		err = json.Unmarshal([]byte(s[0]), mes)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: mes.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "message"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := fmt.Sprintf(key.Message, oid, tid, uid)
		s := mid

		err = d.redigo.Sorted().Delete().Score(k, s)
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
