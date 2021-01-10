package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, obj *timeline.UpdateI) (*timeline.UpdateO, error) {
	{
		ok, err := h.storage.Timeline.Updater.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Timeline.Updater.Update(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
