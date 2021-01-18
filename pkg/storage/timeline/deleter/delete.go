package deleter

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
)

// Delete provides a storage primitive to remove timelines associated with an
// audience.
func (c *Deleter) Delete(req *timeline.DeleteI) (*timeline.DeleteO, error) {
	var err error

	var aid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]
	}

	var tid float64
	{
		tid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.TimelineID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		k := fmt.Sprintf(key.Timeline, aid)
		s := tid

		err = c.redigo.Sorted().Delete().Score(k, s)
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