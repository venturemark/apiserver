package user

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/userid"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	{
		if len(req.Obj) == 0 {
			req.Obj = append(req.Obj, &user.SearchI_Obj{})
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

		if usi == "" && req.Obj[0].Metadata[metadata.TimelineID] == "" && req.Obj[0].Metadata[metadata.VentureID] == "" {
			// User is searching for themselves, but the user ID doesn't exist so the user must not exist, exit early
			return &user.SearchO{}, nil
		}

		if usi != "" {
			for i := range req.Obj {
				if req.Obj[i].Metadata[metadata.UserID] == "" {
					req.Obj[i].Metadata[metadata.UserID] = usi
				} else if req.Obj[i].Metadata[metadata.UserID] != usi {
					return nil, tracer.Mask(invalidInputError)
				}
			}
		}
	}

	{
		ok, err := h.storage.User.Searcher.Verify(ctx, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.User.Searcher.Search(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
