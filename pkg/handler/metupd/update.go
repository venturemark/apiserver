package metupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, obj *metupd.UpdateI) (*metupd.UpdateO, error) {
	{
		ok, err := h.storage.MetUpd.Updater.Verify(obj)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ok {
			res, err := h.storage.MetUpd.Updater.Update(obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			return res, nil
		}
	}

	return nil, tracer.Mask(invalidInputError)
}
