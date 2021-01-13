package message

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, obj *message.SearchI) (*message.SearchO, error) {
	{
		ok, err := h.storage.Message.Searcher.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Message.Searcher.Search(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
