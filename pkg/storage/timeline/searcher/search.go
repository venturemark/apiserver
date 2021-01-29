package searcher

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/schema"
)

// Search provides a filter primitive to lookup timelines associated with a
// user.
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

	var oid string
	{
		oid = req.Obj[0].Metadata[metadata.OrganizationID]
	}

	// With redis we use ZREVRANGE which allows us to search for objects while
	// having support for chunking.
	var str []string
	{
		k := fmt.Sprintf(key.Timeline, oid)
		str, err = s.redigo.Sorted().Search().Index(k, 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return str, nil
}

func (s *Searcher) searchUsr(req *timeline.SearchI) ([]string, error) {
	var err error

	var oid string
	{
		oid = req.Obj[0].Metadata[metadata.OrganizationID]
	}

	var uid string
	{
		uid = req.Obj[0].Metadata[metadata.UserID]
	}

	var aud *audience.SearchO
	{
		aud, err = s.searchAud(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var tim []string
	{
		for _, o := range aud.Obj {
			if !contains(o.Property.User, uid) {
				continue
			}

			tim = append(tim, o.Property.Tmln...)
		}
	}

	var res []string
	{
		for _, t := range tim {
			var tid float64
			{
				tid, err = strconv.ParseFloat(t, 64)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			k := fmt.Sprintf(key.Timeline, oid)
			str, err := s.redigo.Sorted().Search().Score(k, tid, tid)
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
