package message

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, obj *message.CreateI) (*message.CreateO, error) {
	{
		ok, err := h.storage.Message.Creator.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Message.Creator.Create(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
