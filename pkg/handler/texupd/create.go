package texupd

import (
	"context"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/hash"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/user"
)

func (h *Handler) Create(ctx context.Context, req *texupd.CreateI) (*texupd.CreateO, error) {
	{
		u, ok := user.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		req.Obj.Metadata[metadata.UserID] = u
	}

	{
		req.Obj.Metadata[metadata.UpdateID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	}

	{
		rol := &role.CreateI{
			Obj: []*role.CreateI_Obj{
				{
					Metadata: map[string]string{
						metadata.ResourceID:   hash.Update(req.Obj.Metadata),
						metadata.ResourceKind: "update",
						metadata.RoleID:       strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
						metadata.RoleKind:     "owner",
						metadata.SubjectID:    req.Obj.Metadata[metadata.UserID],
						metadata.TimelineID:   req.Obj.Metadata[metadata.TimelineID],
						metadata.UpdateID:     req.Obj.Metadata[metadata.UpdateID],
						metadata.VentureID:    req.Obj.Metadata[metadata.VentureID],
					},
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
		ok, err := h.storage.TexUpd.Creator.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.TexUpd.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
