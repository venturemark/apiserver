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
		ok := h.storage.Metric.Search.Any.Update.Verify(obj)
		if ok {
			res, err := h.storage.Metric.Search.Any.Update.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
