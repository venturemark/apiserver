package deleter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

// Delete provides a storage primitive to remove timelines associated with an
// audience.
func (d *Deleter) Delete(req *timeline.DeleteI) (*timeline.DeleteO, error) {
	var err error

	var tii float64
	{
		tii, err = strconv.ParseFloat(req.Obj.Metadata[metadata.TimelineID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var vei string
	{
		vei = req.Obj.Metadata[metadata.VentureID]
	}

	var tim *schema.Timeline
	{
		k := fmt.Sprintf(key.Timeline, vei)
		s, err := d.redigo.Sorted().Search().Score(k, tii, tii)
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
		k := fmt.Sprintf(key.Timeline, vei)
		s := tii

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *timeline.DeleteO
	{
		res = &timeline.DeleteO{
			Obj: &timeline.DeleteO_Obj{
				Metadata: map[string]string{
					metadata.TimelineStatus: "deleted",
				},
			},
		}
	}

	return res, nil
}
