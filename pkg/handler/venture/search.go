package venture

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/userid"
)

func (h *Handler) Search(ctx context.Context, req *venture.SearchI) (*venture.SearchO, error) {
	{
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
			{
				req.Obj[i].Metadata[metadata.UserID] = usi
			}
		}
	}

	{
		ok, err := h.storage.Venture.Searcher.Verify(ctx, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Venture.Searcher.Search(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
