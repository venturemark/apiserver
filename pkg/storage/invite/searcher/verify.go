package searcher

import (
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Verify(req *invite.SearchI) (bool, error) {
	for _, v := range s.verify {
		ok, err := v.Verify(req)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if !ok {
			return false, nil
		}
	}

	return true, nil
}
