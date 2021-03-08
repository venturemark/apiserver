package audience

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/user"
)

func (h *Handler) Delete(ctx context.Context, req *audience.DeleteI) (*audience.DeleteO, error) {
	{
		u, ok := user.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		for i := range req.Obj {
			req.Obj[i].Metadata[metadata.UserID] = u
		}
	}

	{
		ok, err := h.storage.Audience.Deleter.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Audience.Deleter.Delete(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
