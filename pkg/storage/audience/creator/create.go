package creator

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/index"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *audience.CreateI) (*audience.CreateO, error) {
	var err error

	var aui float64
	{
		aui, err = strconv.ParseFloat(req.Obj.Metadata[metadata.AudienceID], 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var val string
	{
		aud := schema.Audience{
			Obj: schema.AudienceObj{
				Metadata: req.Obj.Metadata,
				Property: schema.AudienceObjProperty{
					Name: req.Obj.Property.Name,
					Tmln: req.Obj.Property.Tmln,
					User: req.Obj.Property.User,
				},
			},
		}

		byt, err := json.Marshal(aud)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(byt)
	}

	{
		k := fmt.Sprintf(key.Audience)
		v := val
		s := aui
		i := index.New(index.Name, req.Obj.Property.Name)

		err = c.redigo.Sorted().Create().Element(k, v, s, i)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *audience.CreateO
	{
		res = &audience.CreateO{
			Obj: &audience.CreateO_Obj{
				Metadata: map[string]string{
					metadata.AudienceID: req.Obj.Metadata[metadata.AudienceID],
				},
			},
		}
	}

	return res, nil
}
