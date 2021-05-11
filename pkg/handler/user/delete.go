package user

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/userid"
)

func (h *Handler) Delete(ctx context.Context, req *user.DeleteI) (*user.DeleteO, error) {
	{
		if len(req.Obj) == 0 {
			req.Obj = append(req.Obj, &user.DeleteI_Obj{})
		}

		for i := range req.Obj {
			if req.Obj[i].Metadata == nil {
				req.Obj[i].Metadata = map[string]string{}
			}
		}
	}

	{
		usi, ok := userid.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		for i := range req.Obj {
			if req.Obj[i].Metadata[metadata.UserID] == "" {
				req.Obj[i].Metadata[metadata.UserID] = usi
			} else if req.Obj[i].Metadata[metadata.UserID] != usi {
				return nil, tracer.Mask(invalidInputError)
			}
		}
	}

	{
		ok, err := h.storage.User.Deleter.Verify(ctx, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.User.Deleter.Delete(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
