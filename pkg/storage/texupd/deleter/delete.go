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

	var tii string
	{
		tii = req.Obj[0].Metadata[metadata.TimelineID]
	}

	var upi float64
	{
		upi, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.UpdateID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var vei string
	{
		vei = req.Obj[0].Metadata[metadata.VentureID]
	}

	var upd *schema.Update
	{
		k := fmt.Sprintf(key.Update, vei, tii)
		s, err := d.redigo.Sorted().Search().Score(k, upi, upi)
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
		k := fmt.Sprintf(key.Update, vei, tii)
		s := upi

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.DeleteO
	{
		res = &texupd.DeleteO{
			Obj: []*texupd.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.UpdateStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
