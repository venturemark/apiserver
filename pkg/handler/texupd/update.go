package texupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, obj *texupd.UpdateI) (*texupd.UpdateO, error) {
	{
		ok, err := h.storage.TexUpd.Updater.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.TexUpd.Updater.Update(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
