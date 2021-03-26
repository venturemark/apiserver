package deleter

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/rescue/pkg/task"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Delete(req *venture.DeleteI) (*venture.DeleteO, error) {
	var err error

	var vek *key.Key
	{
		vek = key.Venture(req.Obj[0].Metadata)
	}

	var ven *schema.Venture
	{
		k := vek.Elem()

		str, err := d.redigo.Simple().Search().Value(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		ven = &schema.Venture{}
		err = json.Unmarshal([]byte(str), ven)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		t := &task.Task{
			Obj: task.TaskObj{
				Metadata: ven.Obj.Metadata,
			},
		}

		t.Obj.Metadata[metadata.TaskAction] = "delete"
		t.Obj.Metadata[metadata.TaskResource] = "venture"

		err = d.rescue.Create(t)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := vek.Elem()

		err = d.redigo.Simple().Delete().Element(k)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *venture.DeleteO
	{
		res = &venture.DeleteO{
			Obj: []*venture.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.VentureID:     vek.ID().S(),
						metadata.VentureStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
