package deleter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

// Delete provides a storage primitive to remove text updates associated with a
// timeline.
func (d *Deleter) Delete(req *texupd.DeleteI) (*texupd.DeleteO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	var uid float64
	{
		uid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.UpdateID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd *schema.Update
	{
		k := fmt.Sprintf(key.UpdateResource, oid, tid)
		s, err := d.redigo.Sorted().Search().Score(k, uid, uid)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		upd = &schema.Update{}
		err = json.Unmarshal([]byte(s[0]), upd)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: upd.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "update"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := fmt.Sprintf(key.UpdateResource, oid, tid)
		s := uid

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.DeleteO
	{
		res = &texupd.DeleteO{
			Obj: &texupd.DeleteO_Obj{
				Metadata: map[string]string{
					metadata.UpdateStatus: "deleted",
				},
			},
		}
	}

	return res, nil
}
