package updater

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"
)

// Update provides a storage primitive to modify text updates associated with a
// timeline. A timeline refers to many updates. For more information about
// technical details see the inline documentation.
func (u *Updater) Update(req *texupd.UpdateI) (*texupd.UpdateO, error) {
	var err error

	var tid string
	{
		tid = req.Obj.Metadata[metadata.TimelineID]
	}

	var uid float64
	{
		uid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.UpdateID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var vid string
	{
		vid = req.Obj.Metadata[metadata.VentureID]
	}

	var upd *schema.Update
	{
		k := fmt.Sprintf(key.Update, vid, tid)
		s, err := u.redigo.Sorted().Search().Score(k, uid, uid)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		upd = &schema.Update{}
		err = json.Unmarshal([]byte(s[0]), upd)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		upd.Obj.Property.Text = req.Obj.Property.Text
	}

	var val string
	{
		byt, err := json.Marshal(upd)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	var mod bool
	{
		k := fmt.Sprintf(key.Update, vid, tid)
		v := val
		s := uid

		mod, err = u.redigo.Sorted().Update().Value(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.UpdateO
	{
		res = &texupd.UpdateO{
			Obj: &texupd.UpdateO_Obj{
				Metadata: map[string]string{},
			},
		}

		if mod {
			res.Obj.Metadata[metadata.UpdateStatus] = "updated"
		}
	}

	return res, nil
}
