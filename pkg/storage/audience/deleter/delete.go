package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *audience.DeleteI) (*audience.DeleteO, error) {
	var err error

	var auk *key.Key
	{
		auk = key.Audience(req.Obj[0].Metadata)
	}

	var aud *schema.Audience
	{
		k := auk.List()
		s, err := d.redigo.Sorted().Search().Score(k, auk.ID().F(), auk.ID().F())
		if err != nil {
			return nil, tracer.Mask(err)
		}

		aud = &schema.Audience{}
		err = json.Unmarshal([]byte(s[0]), aud)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: aud.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "audience"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := auk.List()
		s := auk.ID().F()

		err = d.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.DeleteO
	{
		res = &audience.DeleteO{
			Obj: []*audience.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.AudienceStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
