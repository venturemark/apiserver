package invite

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/userid"
)

func (h *Handler) Update(ctx context.Context, req *invite.UpdateI) (*invite.UpdateO, error) {
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
		ok, err := h.storage.Invite.Updater.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Invite.Updater.Update(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}