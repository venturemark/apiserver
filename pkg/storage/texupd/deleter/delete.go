package deleter

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
)

// Delete provides a storage primitive to remove text updates associated with a
// timeline.
func (c *Deleter) Delete(req *texupd.DeleteI) (*texupd.DeleteO, error) {
	var err error

	var aid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]
	}

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

	{
		k := fmt.Sprintf(key.Update, aid, tid)
		s := uid

		err = c.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *texupd.DeleteO
	{
		res = &texupd.DeleteO{
			Obj: &texupd.DeleteO_Obj{
				Metadata: map[string]string{
					metadata.UpdateStatus: "deleted",
				},
			},
		}
	}

	return res, nil
}
