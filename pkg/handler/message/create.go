package message

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/user"
)

func (h *Handler) Create(ctx context.Context, req *message.CreateI) (*message.CreateO, error) {
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
						metadata.ResourceKind: "message",
						metadata.RoleKind:     "owner",
						metadata.SubjectID:    req.Obj.Metadata[metadata.UserID],
						metadata.TimelineID:   req.Obj.Metadata[metadata.TimelineID],
						metadata.UpdateID:     req.Obj.Metadata[metadata.UpdateID],
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
		ok, err := h.storage.Message.Creator.Verify(req)
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
		res, err := h.storage.Message.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
