package audience

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, obj *audience.UpdateI) (*audience.UpdateO, error) {
	{
		ok, err := h.storage.Audience.Updater.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Audience.Updater.Update(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
