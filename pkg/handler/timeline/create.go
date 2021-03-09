package timeline

import (
	"context"
	"strconv"
	"time"

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

		for i := range req.Obj {
			req.Obj[i].Metadata[metadata.UserID] = u
		}
	}

	{
		for i := range req.Obj {
			req.Obj[i].Metadata[metadata.TimelineID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		}
	}

	{
		rol := &role.CreateI{
			Obj: []*role.CreateI_Obj{
				{
					Metadata: map[string]string{
						metadata.ResourceKind: "timeline",
						metadata.RoleID:       strconv.FormatInt(time.Now().UTC().UnixNano(), 10),
						metadata.RoleKind:     "owner",
						metadata.SubjectID:    req.Obj[0].Metadata[metadata.UserID],
						metadata.TimelineID:   req.Obj[0].Metadata[metadata.TimelineID],
						metadata.VentureID:    req.Obj[0].Metadata[metadata.VentureID],
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
		ok, err := h.storage.Timeline.Creator.Verify(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err := h.storage.Timeline.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		return res, nil
	}
}
