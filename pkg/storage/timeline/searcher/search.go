package searcher

import (
	"encoding/json"
	"fmt"
	"github.com/venturemark/permission/pkg/label/visibility"
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
		venEmp := req.Obj[0].Metadata[metadata.VentureID] == ""

		if subEmp && timEmp { // search by venture id
			str, err = s.searchAll(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else if !subEmp && timEmp && venEmp { // search by subject id in all ventures
			str, err = s.searchSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else if !subEmp && timEmp && !venEmp { // search by subject id in a specific venture
			str, err = s.searchVenSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else if subEmp && !timEmp { // search by subject id and timeline id
			str, err = s.searchTimSub(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	var res *timeline.SearchO
	{
		res = &timeline.SearchO{}

		for _, timelineString := range str {
			tim := &schema.Timeline{}
			err := json.Unmarshal([]byte(timelineString), tim)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			{
				lastUpdate, err := s.searchLastUpdate(tim)
				if err != nil {
					return nil, tracer.Mask(err)
				}

				if lastUpdate != "" {
					update := schema.Update{}
					err := json.Unmarshal([]byte(lastUpdate), &update)
					if err != nil {
						return nil, tracer.Mask(err)
					}

					updateID := update.Obj.Metadata[metadata.UpdateID]
					tim.Obj.Metadata[metadata.TimelineLastUpdate] = updateID
				}
			}

			if req.Obj[0].Metadata[metadata.UserID] == "" && tim.Obj.Metadata[metadata.ResourceVisibility] != visibility.Public.Label() {
				// Only return public timelines on unauthenticated requests
				continue
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
	timSub, err := s.searchTimSub(req)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	venSub, err := s.searchVenSub(req)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	combined := timSub[:]
	for _, l := range venSub {
		if !contains(combined, l) {
			combined = append(combined, l)
		}
	}

	return combined, nil
}

func (s *Searcher) searchVenSub(req *timeline.SearchI) ([]string, error) {
	var str []string

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

func (s *Searcher) searchTimSub(req *timeline.SearchI) ([]string, error) {
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

func (s *Searcher) searchLastUpdate(tim *schema.Timeline) (string, error) {
	var err error

	var upk *key.Key
	{
		upk = key.Update(tim.Obj.Metadata)
	}

	var str []string
	{
		k := upk.List()

		str, err = s.redigo.Sorted().Search().Order(k, 0, 1)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	if len(str) == 0 {
		return "", nil
	}

	return str[0], nil
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
