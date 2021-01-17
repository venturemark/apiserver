package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, obj *timeline.DeleteI) (*timeline.DeleteO, error) {
	{
		ok, err := h.storage.Timeline.Deleter.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Timeline.Deleter.Delete(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
