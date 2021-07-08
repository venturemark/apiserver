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
	var err error

	var str []string
	{
		resEmp := req.Obj[0].Metadata[metadata.ResourceKind] == ""

		if !resEmp {
			str, err = s.searchRes(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if resEmp {
			str, err = s.searchSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	usi := req.Obj[0].Metadata[metadata.UserID]

	var res *user.SearchO
	{
		res = &user.SearchO{}

		for _, s := range str {
			use := &schema.User{}
			err := json.Unmarshal([]byte(s), use)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			// Don't leak private info about other users
			if use.Obj.Metadata[metadata.UserID] != usi {
				delete(use.Obj.Metadata, metadata.TimelineLastUpdate)
			}

			o := &user.SearchO_Obj{
				Metadata: use.Obj.Metadata,
				Property: &user.SearchO_Obj_Property{
					Desc: use.Obj.Property.Desc,
					Name: use.Obj.Property.Name,
					Mail: use.Obj.Property.Mail,
					Prof: prof(use.Obj.Property.Prof),
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}

func (s *Searcher) searchRes(req *user.SearchI) ([]string, error) {
	var err error

	var rol []*schema.Role
	{
		rol, err = s.searchRol(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var str []string
	{
		for _, r := range rol {
			{
				r.Obj.Metadata[metadata.UserID] = r.Obj.Metadata[metadata.SubjectID]
			}

			req := &user.SearchI{
				Obj: []*user.SearchI_Obj{
					{
						Metadata: r.Obj.Metadata,
					},
				},
			}

			lis, err := s.searchSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			str = append(str, lis...)
		}
	}

	return str, nil
}

func (s *Searcher) searchRol(req *user.SearchI) ([]*schema.Role, error) {
	var err error

	var rok *key.Key
	{
		rok = key.Role(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := rok.List()

		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var rol []*schema.Role
	{
		for _, s := range str {
			r := &schema.Role{}
			err := json.Unmarshal([]byte(s), r)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			rol = append(rol, r)
		}
	}

	return rol, nil
}

func (s *Searcher) searchSub(req *user.SearchI) ([]string, error) {
	var usk *key.Key
	{
		sub := req.Obj[0].Metadata[metadata.SubjectID] != ""
		use := req.Obj[0].Metadata[metadata.UserID] != ""

		usk = key.User(req.Obj[0].Metadata)
		if sub && !use {
			usk = key.Subject(req.Obj[0].Metadata)
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

	return str, nil
}

func prof(i []schema.UserObjPropertyProf) []*user.SearchO_Obj_Property_Prof {
	var o []*user.SearchO_Obj_Property_Prof

	for _, l := range i {
		o = append(o, &user.SearchO_Obj_Property_Prof{Desc: l.Desc, Vent: l.Vent})
	}

	return o
}
