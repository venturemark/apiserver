package role

import (
	"context"

	"github.com/venturemark/apicommon/pkg/hash"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/user"
)

func (h *Handler) Search(ctx context.Context, req *role.SearchI) (*role.SearchO, error) {
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
		for i := range req.Obj {
			switch req.Obj[i].Metadata[metadata.ResourceKind] {
			case "audience":
				req.Obj[i].Metadata[metadata.ResourceID] = hash.Audience(req.Obj[i].Metadata)
			case "message":
				req.Obj[i].Metadata[metadata.ResourceID] = hash.Message(req.Obj[i].Metadata)
			case "timeline":
				req.Obj[i].Metadata[metadata.ResourceID] = hash.Timeline(req.Obj[i].Metadata)
			case "update":
				req.Obj[i].Metadata[metadata.ResourceID] = hash.Update(req.Obj[i].Metadata)
			case "venture":
				req.Obj[i].Metadata[metadata.ResourceID] = hash.Venture(req.Obj[i].Metadata)
			}
		}
	}

	{
		ok, err := h.storage.Role.Searcher.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Role.Searcher.Search(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
