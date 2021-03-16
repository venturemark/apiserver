package audience

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/venturemark/apiserver/pkg/context/userid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *audience.SearchI) (*audience.SearchO, error) {
	{
		usi, ok := userid.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		for i := range req.Obj {
			req.Obj[i].Metadata[metadata.UserID] = usi
		}
	}

	{
		ok, err := h.storage.Audience.Searcher.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Audience.Searcher.Search(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
