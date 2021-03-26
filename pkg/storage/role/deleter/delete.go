package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *role.DeleteI) (*role.DeleteO, error) {
	var err error

	var rok *key.Key
	{
		rok = key.Role(req.Obj[0].Metadata)
	}

	var rol *schema.Role
	{
		k := rok.List()
		s := rok.ID().F()

		str, err := d.redigo.Sorted().Search().Score(k, s, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		rol = &schema.Role{}
		err = json.Unmarshal([]byte(str[0]), rol)
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
		k := rok.List()
		s := rok.ID().F()

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
						metadata.RoleID:     rok.ID().S(),
						metadata.RoleStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
