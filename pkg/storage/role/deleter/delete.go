package deleter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *role.DeleteI) (*role.DeleteO, error) {
	var err error

	var rid float64
	{
		rid, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.RoleID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sid string
	{
		sid = req.Obj[0].Metadata[metadata.SubjectID]
	}

	var rol *schema.Role
	{
		k := fmt.Sprintf(key.Role, sid)
		s, err := d.redigo.Sorted().Search().Score(k, rid, rid)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		rol = &schema.Role{}
		err = json.Unmarshal([]byte(s[0]), rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: rol.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "role"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := fmt.Sprintf(key.Role, sid)
		s := rid

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *role.DeleteO
	{
		res = &role.DeleteO{
			Obj: []*role.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.RoleStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
