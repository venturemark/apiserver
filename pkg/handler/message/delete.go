package message

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, obj *message.DeleteI) (*message.DeleteO, error) {
	{
		ok, err := h.storage.Message.Deleter.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.Message.Deleter.Delete(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
