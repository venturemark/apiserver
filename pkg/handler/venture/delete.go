package venture

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/user"
)

func (h *Handler) Delete(ctx context.Context, req *venture.DeleteI) (*venture.DeleteO, error) {
	{
		for i := range req.Obj {
			if req.Obj[i].Metadata == nil {
				req.Obj[i].Metadata = map[string]string{}
			}
		}
	}

	{
		u, ok := user.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		for i := range req.Obj {
			{
				req.Obj[i].Metadata[metadata.UserID] = u
			}
		}
	}

	{
		ok, err := h.storage.Venture.Deleter.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Venture.Deleter.Delete(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
