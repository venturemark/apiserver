package creator

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"
)

func (c *Creator) Verify(ctx context.Context, req *user.CreateI) (bool, error) {
	for _, v := range c.verify {
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
