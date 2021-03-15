package searcher

import (
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *user.SearchI) (*user.SearchO, error) {
	var usk *key.Key
	{
		sub := req.Obj[0].Metadata[metadata.SubjectID] != ""
		use := req.Obj[0].Metadata[metadata.UserID] != ""

		if sub && !use {
			usk = key.Subject(req.Obj[0].Metadata)
		}

		if !sub && use {
			usk = key.User(req.Obj[0].Metadata)
		}
	}

	var str []string
	{
		k := usk.Elem()

		s, err := s.redigo.Simple().Search().Value(k)
		if simple.IsNotFound(err) {
			// fall through
		} else if err != nil {
			return nil, tracer.Mask(err)
		} else {
			str = append(str, s)
		}
	}

	var res *user.SearchO
	{
		res = &user.SearchO{}

		for _, s := range str {
			use := &schema.User{}
			err := json.Unmarshal([]byte(s), use)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &user.SearchO_Obj{
				Metadata: use.Obj.Metadata,
				Property: &user.SearchO_Obj_Property{
					Desc: use.Obj.Property.Desc,
					Name: use.Obj.Property.Name,
					Prof: prof(use.Obj.Property.Prof),
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}

func prof(i []schema.UserObjPropertyProf) []*user.SearchO_Obj_Property_Prof {
	var o []*user.SearchO_Obj_Property_Prof

	for _, l := range i {
		o = append(o, &user.SearchO_Obj_Property_Prof{Desc: l.Desc, Vent: l.Vent})
	}

	return o
}
