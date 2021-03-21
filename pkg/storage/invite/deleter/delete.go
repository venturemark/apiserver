package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *invite.DeleteI) (*invite.DeleteO, error) {
	var err error

	var ink *key.Key
	{
		ink = key.Invite(req.Obj[0].Metadata)
	}

	var inv *schema.Invite
	{
		k := ink.List()
		s, err := d.redigo.Sorted().Search().Score(k, ink.ID().F(), ink.ID().F())
		if err != nil {
			return nil, tracer.Mask(err)
		}

		inv = &schema.Invite{}
		err = json.Unmarshal([]byte(s[0]), inv)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: inv.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "invite"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := ink.List()
		s := ink.ID().F()

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *invite.DeleteO
	{
		res = &invite.DeleteO{
			Obj: []*invite.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.InviteStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
