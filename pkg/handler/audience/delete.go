package audience

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, obj *audience.DeleteI) (*audience.DeleteO, error) {
	{
		ok, err := h.storage.Audience.Deleter.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Audience.Deleter.Delete(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
