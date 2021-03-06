package timeline

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/user"
)

func (h *Handler) Create(ctx context.Context, req *timeline.CreateI) (*timeline.CreateO, error) {
	{
		u, ok := user.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		req.Obj.Metadata[metadata.UserID] = u
	}

	var rol *role.CreateI
	{
		rol = &role.CreateI{
			Obj: []*role.CreateI_Obj{
				{
					Metadata: map[string]string{
						metadata.ResourceKind: "timeline",
						metadata.RoleKind:     "owner",
						metadata.SubjectID:    req.Obj.Metadata[metadata.UserID],
						metadata.VentureID:    req.Obj.Metadata[metadata.VentureID],
					},
				},
			},
		}
	}

	{
		ok, err := h.storage.Role.Creator.Verify(rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}
	}

	{
		ok, err := h.storage.Timeline.Creator.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}
	}

	{
		_, err := h.storage.Role.Creator.Create(rol)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		res, err := h.storage.Timeline.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
