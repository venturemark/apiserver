package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *user.DeleteI) (*user.DeleteO, error) {
	var err error

	var usk *key.Key
	{
		usk = key.User(req.Obj[0].Metadata)
	}

	var use *schema.User
	{
		k := usk.Elem()

		str, err := d.redigo.Simple().Search().Value(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		use = &schema.User{}
		err = json.Unmarshal([]byte(str), use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: use.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "user"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := usk.Elem()

		err = d.redigo.Simple().Delete().Element(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *user.DeleteO
	{
		res = &user.DeleteO{
			Obj: []*user.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.UserID:     usk.ID().S(),
						metadata.UserStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
