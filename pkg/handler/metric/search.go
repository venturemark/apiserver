package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, obj *metric.SearchI) (*metric.SearchO, error) {
	{
		ok, err := h.storage.Metric.Searcher.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Metric.Searcher.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
