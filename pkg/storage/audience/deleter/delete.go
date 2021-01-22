package deleter

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
)

func (c *Deleter) Delete(req *audience.DeleteI) (*audience.DeleteO, error) {
	var err error

	var aid float64
	{
		aid, err = strconv.ParseFloat(req.Obj.Metadata[metadata.AudienceID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var oid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
	}

	{
		k := fmt.Sprintf(key.Audience, oid)
		s := aid

		err = c.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.DeleteO
	{
		res = &audience.DeleteO{
			Obj: &audience.DeleteO_Obj{
				Metadata: map[string]string{
					metadata.AudienceStatus: "deleted",
				},
			},
		}
	}

	return res, nil
}
