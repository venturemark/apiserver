package searcher

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (s *Searcher) Verify(ctx context.Context, req *role.SearchI) (bool, error) {
	for _, v := range s.verify {
		ok, err := v.Verify(ctx, req)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if !ok {
			return false, nil
		}
	}

	return true, nil
}
