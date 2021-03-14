package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *timeline.DeleteI) (*timeline.DeleteO, error) {
	var err error

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	var tim *schema.Timeline
	{
		k := tik.List()
		s, err := d.redigo.Sorted().Search().Score(k, tik.ID().F(), tik.ID().F())
		if err != nil {
			return nil, tracer.Mask(err)
		}

		tim = &schema.Timeline{}
		err = json.Unmarshal([]byte(s[0]), tim)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: tim.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "timeline"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := tik.List()
		s := tik.ID().F()

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.DeleteO
	{
		res = &timeline.DeleteO{
			Obj: []*timeline.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.TimelineStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
