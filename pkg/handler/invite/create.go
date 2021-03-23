package invite

import (
	"context"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/userid"
	"github.com/venturemark/apiserver/pkg/random"
)

func (h *Handler) Create(ctx context.Context, req *invite.CreateI) (*invite.CreateO, error) {
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
				req.Obj[i].Metadata[metadata.SubjectID] = usi
				req.Obj[i].Metadata[metadata.UserID] = usi
			}

			{
				req.Obj[i].Metadata[metadata.ResourceKind] = "invite"
				req.Obj[i].Metadata[metadata.RoleKind] = "owner"
			}

			{
				req.Obj[i].Metadata[metadata.InviteCode] = random.MustNew()
			}

			{
				req.Obj[i].Metadata[metadata.RoleID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
				req.Obj[i].Metadata[metadata.InviteID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			}
		}
	}

	{
		rol := &role.CreateI{
			Obj: []*role.CreateI_Obj{
				{
					Metadata: req.Obj[0].Metadata,
				},
			},
		}

		ok, err := h.storage.Role.Creator.Verify(rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		_, err = h.storage.Role.Creator.Create(rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		ok, err := h.storage.Invite.Creator.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Invite.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
