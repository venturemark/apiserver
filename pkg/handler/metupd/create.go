package metupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, obj *metupd.CreateI) (*metupd.CreateO, error) {
	// Create metric updates associated with the given timeline.
	{
		ok := h.storage.MetUpd.Create.Verify(obj)
		if ok {
			res, err := h.storage.MetUpd.Create.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
