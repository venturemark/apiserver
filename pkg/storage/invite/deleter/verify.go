package deleter

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/tracer"
)

func (d *Deleter) Verify(ctx context.Context, req *invite.DeleteI) (bool, error) {
	for _, v := range d.verify {
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
