package update

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/update"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, obj *update.SearchI) (*update.SearchO, error) {
	{
		ok, err := h.storage.Update.Searcher.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Update.Searcher.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
