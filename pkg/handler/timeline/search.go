package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, obj *timeline.SearchI) (*timeline.SearchO, error) {
	// Search for any timeline associated with the given user. One user ID must
	// be provided.
	{
		ok, err := h.storage.Timeline.Search.Non.User.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Timeline.Search.Non.User.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}