package searcher

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

const (
	timPat = "ven:*:tim:sub:%s"
	venPat = "ven:sub:%s"
)

func (s *Searcher) Search(req *timeline.SearchI) (*timeline.SearchO, error) {
	var err error

	var str []string
	{
		subEmp := req.Obj[0].Metadata[metadata.SubjectID] == ""
		timEmp := req.Obj[0].Metadata[metadata.TimelineID] == ""

		if subEmp && timEmp {
			str, err = s.searchAll(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if !subEmp && timEmp {
			str, err = s.searchSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if subEmp && !timEmp {
			str, err = s.searchTim(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	var res *timeline.SearchO
	{
		res = &timeline.SearchO{}

		for _, s := range str {
			tim := &schema.Timeline{}
			err := json.Unmarshal([]byte(s), tim)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			o := &timeline.SearchO_Obj{
				Metadata: tim.Obj.Metadata,
				Property: &timeline.SearchO_Obj_Property{
					Desc: tim.Obj.Property.Desc,
					Name: tim.Obj.Property.Name,
					Stat: tim.Obj.Property.Stat,
				},
			}

			res.Obj = append(res.Obj, o)
		}
	}

	return res, nil
}

func (s *Searcher) searchAll(req *timeline.SearchI) ([]string, error) {
	var err error

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := tik.List()

		str, err = s.redigo.Sorted().Search().Order(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return str, nil
}

func (s *Searcher) searchRolPat(pat string, req *timeline.SearchI) ([]*schema.Role, error) {
	var err error

	var sui string
	{
		sui = req.Obj[0].Metadata[metadata.SubjectID]
	}

	var don chan struct{}
	var erc chan error
	var res chan string
	{
		don = make(chan struct{}, 1)
		erc = make(chan error, 1)
		res = make(chan string, 1)
	}

	var rol []*schema.Role
	go func() {
		defer close(don)

		for k := range res {
			str, err := s.redigo.Sorted().Search().Order(k, 0, -1)
			if err != nil {
				erc <- tracer.Mask(err)
			}

			for _, k := range str {
				rei, roi := split(k)

				val, err := s.redigo.Sorted().Search().Score(rei, roi, roi)
				if err != nil {
					erc <- tracer.Mask(err)
				}

				if len(val) == 0 {
					continue
				}

				r := &schema.Role{}
				err = json.Unmarshal([]byte(val[0]), r)
				if err != nil {
					erc <- tracer.Mask(err)
				}

				rol = append(rol, r)
			}
		}
	}()

	go func() {
		defer close(res)

		k := fmt.Sprintf(pat, sui)

		err = s.redigo.Walker().Simple(k, don, res)
		if err != nil {
			erc <- tracer.Mask(err)
		}
	}()

	{
		select {
		case <-don:
			return rol, nil

		case err := <-erc:
			return nil, tracer.Mask(err)

		case <-time.After(1 * time.Second):
			return nil, tracer.Mask(timeoutError)
		}
	}
}

func (s *Searcher) searchSub(req *timeline.SearchI) ([]string, error) {
	var str []string

	{
		rol, err := s.searchRolPat(timPat, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		for _, r := range rol {
			req := &timeline.SearchI{
				Obj: []*timeline.SearchI_Obj{
					{
						Metadata: r.Obj.Metadata,
					},
				},
			}

			lis, err := s.searchTim(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			for _, l := range lis {
				if !contains(str, l) {
					str = append(str, l)
				}
			}
		}
	}

	{
		rol, err := s.searchRolPat(venPat, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		for _, r := range rol {
			req := &timeline.SearchI{
				Obj: []*timeline.SearchI_Obj{
					{
						Metadata: r.Obj.Metadata,
					},
				},
			}

			lis, err := s.searchAll(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			for _, l := range lis {
				if !contains(str, l) {
					str = append(str, l)
				}
			}
		}
	}

	return str, nil
}

func (s *Searcher) searchTim(req *timeline.SearchI) ([]string, error) {
	var err error

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	var str []string
	{
		k := tik.List()
		c := tik.ID().F()

		str, err = s.redigo.Sorted().Search().Score(k, c, c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return str, nil
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

func contains(l []string, s string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}

	return false
}
