package searcher

import (
	"encoding/json"
	"fmt"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *role.SearchI) (*role.SearchO, error) {
	var err error

	var rei string
	{
		rei = req.Obj[0].Metadata[metadata.ResourceID]
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := fmt.Sprintf(key.Role, rei)
		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *role.SearchO
	{
		res = &role.SearchO{}

		for _, s := range str {
			rol := &schema.Role{}
			err := json.Unmarshal([]byte(s), rol)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &role.SearchO_Obj{
				Metadata: rol.Obj.Metadata,
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}