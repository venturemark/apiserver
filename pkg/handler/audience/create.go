package audience

import (
	"context"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/venturemark/apiserver/pkg/context/userid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *audience.CreateI) (*audience.CreateO, error) {
	{
		u, ok := userid.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		for i := range req.Obj {
			req.Obj[i].Metadata[metadata.UserID] = u
		}
	}

	{
		for i := range req.Obj {
			req.Obj[i].Metadata[metadata.AudienceID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		}
	}

	{
		rol := &role.CreateI{
			Obj: []*role.CreateI_Obj{
				{
					Metadata: map[string]string{
						metadata.AudienceID:   req.Obj[0].Metadata[metadata.AudienceID],
						metadata.ResourceKind: "audience",
						metadata.RoleID:       strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
						metadata.RoleKind:     "owner",
						metadata.SubjectID:    req.Obj[0].Metadata[metadata.UserID],
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
		ok, err := h.storage.Audience.Creator.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Audience.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
