package searcher

import (
	"encoding/json"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Search(req *timeline.SearchI) (*timeline.SearchO, error) {
	var err error

	var str []string
	{
		met := map[string]string{
			metadata.PermissionID:     "audience",
			metadata.PermissionStatus: "enabled",
		}

		con := metadata.Contains(req.Obj[0].Metadata, met)

		if con {
			str, err = s.searchUsr(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		} else {
			str, err = s.searchAll(req)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	// We store timelines in a sorted set. The elements of the sorted set are
	// concatenated strings of the unix timestamp of timeline creation and the
	// timeline name.
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

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
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

func (s *Searcher) searchUsr(req *timeline.SearchI) ([]string, error) {
	var err error

	var aud *audience.SearchO
	{
		aud, err = s.searchAud(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	var usi string
	{
		usi = req.Obj[0].Metadata[metadata.UserID]
	}

	var tim []string
	{
		for _, o := range aud.Obj {
			if !contains(o.Property.User, usi) {
				continue
			}

			tim = append(tim, o.Property.Tmln...)
		}
	}

	var res []string
	{
		for _, t := range tim {
			var tii float64
			{
				tii, err = strconv.ParseFloat(t, 64)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			k := tik.List()
			str, err := s.redigo.Sorted().Search().Score(k, tii, tii)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			res = append(res, str...)
		}
	}

	return res, nil
}

func contains(all []string, sub string) bool {
	for _, a := range all {
		if a == sub {
			return true
		}
	}

	return false
}
