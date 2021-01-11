package audience

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, obj *audience.CreateI) (*audience.CreateO, error) {
	{
		ok, err := h.storage.Audience.Creator.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Audience.Creator.Create(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
