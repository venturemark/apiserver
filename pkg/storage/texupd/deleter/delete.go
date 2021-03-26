package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *texupd.DeleteI) (*texupd.DeleteO, error) {
	var err error

	var upk *key.Key
	{
		upk = key.Update(req.Obj[0].Metadata)
	}

	var upd *schema.Update
	{
		k := upk.List()
		s, err := d.redigo.Sorted().Search().Score(k, upk.ID().F(), upk.ID().F())
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
		k := upk.List()
		s := upk.ID().F()

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
						metadata.UpdateID:     upk.ID().S(),
						metadata.UpdateStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
