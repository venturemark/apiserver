package texupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Delete(ctx context.Context, obj *texupd.DeleteI) (*texupd.DeleteO, error) {
	{
		ok, err := h.storage.TexUpd.Deleter.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.TexUpd.Deleter.Delete(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
