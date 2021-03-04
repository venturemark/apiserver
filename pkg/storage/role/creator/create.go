package creator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/resource"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Create(req *role.CreateI) (*role.CreateO, error) {
	var err error

	// We manage data on a timeline. Our main identifier is a unix timestamp in
	// nano seconds, normalized to the UTC timezone. Our discovery mechanisms is
	// designed based on this very unix timestamp. Everything starts with time,
	// which means that pseudo random IDs are irrelevant for us. Note that we
	// tracked IDs once in seconds, which caused problems when progammatically
	// faking demo timelines, because only one timeline per second could be
	// created.
	var rid float64
	{
		rid = float64(time.Now().UTC().UnixNano())
	}

	var sha string
	{
		sha = resource.Hash(req.Obj[0].Metadata)
	}

	{
		req.Obj[0].Metadata[metadata.RoleID] = strconv.FormatFloat(rid, 'f', -1, 64)
	}

	var res string
	var val string
	{
		rol := schema.Role{
			Obj: schema.RoleObj{
				Metadata: req.Obj[0].Metadata,
				Property: schema.RoleObjProperty{
					Kin: req.Obj[0].Property.Kin,
					Res: req.Obj[0].Property.Res,
				},
			},
		}

		byt, err := json.Marshal(rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		res = rol.Resource()
		val = string(byt)
	}

	{
		k := fmt.Sprintf(key.Role, sha)
		v := val
		s := rid

		err = c.redigo.Sorted().Create().Element(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *role.CreateO
	{
		res = &role.CreateO{
			Obj: []*role.CreateO_Obj{
				{
					Metadata: map[string]string{
						metadata.RoleID: req.Obj[0].Metadata[metadata.RoleID],
					},
				},
			},
		}
	}

	return res, nil
}
