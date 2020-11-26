package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, obj *timeline.CreateI) (*timeline.CreateO, error) {
	{
		ok, err := h.storage.Timeline.Creator.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Timeline.Creator.Create(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
