package searcher

import (
	"encoding/json"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"
	"strconv"
	"strings"
)

func (s *Searcher) Search(req *venture.SearchI) (*venture.SearchO, error) {
	var err error

	var str []string
	{
		suiEmp := req.Obj[0].Metadata[metadata.SubjectID] == ""
		veiEmp := req.Obj[0].Metadata[metadata.VentureID] == ""

		if !suiEmp && veiEmp {
			str, err = s.searchSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if suiEmp && !veiEmp {
			str, err = s.searchVen(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	var res *venture.SearchO
	{
		res = &venture.SearchO{}

		for _, s := range str {
			ven := &schema.Venture{}
			err := json.Unmarshal([]byte(s), ven)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &venture.SearchO_Obj{
				Metadata: ven.Obj.Metadata,
				Property: &venture.SearchO_Obj_Property{
					Desc: ven.Obj.Property.Desc,
					Link: lin(ven.Obj.Property.Link),
					Name: ven.Obj.Property.Name,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}

func (s *Searcher) searchRol(req *venture.SearchI) ([]*schema.Role, error) {
	var err error

	{
		req.Obj[0].Metadata[metadata.ResourceKind] = "venture"
	}

	var suk *key.Key
	{
		suk = key.Subject(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := suk.Elem()

		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var rol []*schema.Role
	{
		for _, k := range str {
			rei, roi := split(k)

			val, err := s.redigo.Sorted().Search().Score(rei, roi, roi)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			if len(val) == 0 {
				continue
			}

			r := &schema.Role{}
			err = json.Unmarshal([]byte(val[0]), r)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			rol = append(rol, r)
		}
	}

	return rol, nil
}

func (s *Searcher) searchSub(req *venture.SearchI) ([]string, error) {
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
			req := &venture.SearchI{
				Obj: []*venture.SearchI_Obj{
					{
						Metadata: r.Obj.Metadata,
					},
				},
			}

			lis, err := s.searchVen(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			str = append(str, lis...)
		}
	}

	return str, nil
}

func (s *Searcher) searchVen(req *venture.SearchI) ([]string, error) {
	var vek *key.Key
	{
		vek = key.Venture(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := vek.Elem()

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

func lin(i []schema.VentureObjPropertyLink) []*venture.SearchO_Obj_Property_Link {
	var o []*venture.SearchO_Obj_Property_Link

	for _, l := range i {
		o = append(o, &venture.SearchO_Obj_Property_Link{Addr: l.Addr, Text: l.Text})
	}

	return o
}

func split(s string) (string, float64) {
	var err error

	i := strings.LastIndex(s, ":")

	var rei string
	{
		rei = s[:i]
	}

	var roi float64
	{
		roi, err = strconv.ParseFloat(s[i+1:], 64)
		if err != nil {
			panic(err)
		}
	}

	return rei, roi
}
