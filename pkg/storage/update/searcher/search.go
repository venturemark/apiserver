package searcher

import (
	"encoding/json"
	"fmt"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/update"
	"github.com/xh3b4sd/tracer"
)

// Search provides a filter primitive to lookup updates associated with a
// timeline. A timeline refers to many updates. Updates can be found considering
// their scope and time of creation. For more information about technical
// details see the inline documentation.
func (s *Searcher) Search(req *update.SearchI) (*update.SearchO, error) {
	var err error

	var oid string
	{
		oid = req.Obj[0].Metadata[metadata.OrganizationID]
	}

	var tid string
	{
		tid = req.Obj[0].Metadata[metadata.TimelineID]
	}

	var str []string
	{
		k := fmt.Sprintf(key.Update, oid, tid)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *update.SearchO
	{
		res = &update.SearchO{}

		for _, s := range str {
			upd := &schema.Update{}
			err := json.Unmarshal([]byte(s), upd)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &update.SearchO_Obj{
				Metadata: upd.Obj.Metadata,
				Property: &update.SearchO_Obj_Property{
					Text: upd.Obj.Property.Text,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}
