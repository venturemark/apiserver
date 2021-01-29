package deleter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/schema"
)

// Delete provides a storage primitive to remove timelines associated with an
// audience.
func (d *Deleter) Delete(req *timeline.DeleteI) (*timeline.DeleteO, error) {
	var err error

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	var tid float64
	{
		tid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.TimelineID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tim *schema.Timeline
	{
		k := fmt.Sprintf(key.Timeline, oid)
		s, err := d.redigo.Sorted().Search().Score(k, tid, tid)
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
		k := fmt.Sprintf(key.Timeline, oid)
		s := tid

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
