package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, obj *metric.SearchI) (*metric.SearchO, error) {
	// Search for any metric associated with the given updates. One or many
	// update IDs may be provided.
	{
		ok, err := h.storage.Metric.Search.Non.Timeline.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Metric.Search.Non.Timeline.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
