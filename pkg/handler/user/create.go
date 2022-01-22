package user

import (
	"context"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/context/claimid"
)

func (h *Handler) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
	{
		for i := range req.Obj {
			if req.Obj[i].Metadata == nil {
				req.Obj[i].Metadata = map[string]string{}
			}
		}
	}

	{
		cli, ok := claimid.FromContext(ctx)
		if !ok {
			return nil, tracer.Mask(invalidUserError)
		}

		for i := range req.Obj {
			{
				req.Obj[i].Metadata[metadata.ResourceKind] = "user"
				req.Obj[i].Metadata[metadata.RoleKind] = "owner"
			}

			{
				req.Obj[i].Metadata[metadata.RoleID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
				req.Obj[i].Metadata[metadata.UserID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			}

			{
				req.Obj[i].Metadata[metadata.ClaimID] = cli
				req.Obj[i].Metadata[metadata.SubjectID] = req.Obj[i].Metadata[metadata.UserID]
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

		ok, err := h.storage.Role.Creator.Verify(ctx, rol)
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

	var res *user.CreateO
	{
		ok, err := h.storage.User.Creator.Verify(ctx, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err = h.storage.User.Creator.Create(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for i := range res.Obj {
		createdUser := res.Obj[i]
		usi := createdUser.Metadata[metadata.UserID]
		userRequest := req.Obj[i]
		err := h.createDefaultTimelines(ctx, userRequest, usi)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return res, nil
}

var defaultVenture = venture.CreateI_Obj_Property{
	Desc: "x",
	Name: "x",
}

var defaultTimelinesMap = map[string][]timeline.CreateI_Obj_Property{
	"a": {
		{
			Desc: "a",
			Name: "a",
		},
	},
	"b": {
		{
			Desc: "b",
			Name: "b",
		},
	},
}

func (h *Handler) createDefaultTimelines(ctx context.Context, userRequest *user.CreateI_Obj, usi string) error {
	defaultTimelines := defaultTimelinesMap[userRequest.Metadata["user.venturemark.co/prepopulate"]]

	if len(defaultTimelines) == 0 {
		return nil
	}

	ventureRoleID := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	ventureID := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	ventureMetadata := map[string]string{
		metadata.SubjectID:    usi,
		metadata.UserID:       usi,
		metadata.ResourceKind: "venture",
		metadata.RoleKind:     "owner",
		metadata.RoleID:       ventureRoleID,
		metadata.VentureID:    ventureID,
	}

	{
		rolReq := role.CreateI{
			Obj: []*role.CreateI_Obj{
				{
					Metadata: ventureMetadata,
				},
			},
		}

		ok, err := h.storage.Role.Creator.Verify(ctx, &rolReq)
		if err != nil {
			return tracer.Mask(err)
		}

		if !ok {
			return tracer.Mask(invalidInputError)
		}

		_, err = h.storage.Role.Creator.Create(&rolReq)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		ventureReq := venture.CreateI{
			Obj: []*venture.CreateI_Obj{
				{
					Metadata: ventureMetadata,
					Property: &venture.CreateI_Obj_Property{
						Desc: defaultVenture.Desc,
						Name: defaultVenture.Name,
					},
				},
			},
		}

		ok, err := h.storage.Venture.Creator.Verify(ctx, &ventureReq)
		if err != nil {
			return tracer.Mask(err)
		}

		if !ok {
			return tracer.Mask(invalidInputError)
		}

		_, err = h.storage.Venture.Creator.Create(&ventureReq)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, defaultTimeline := range defaultTimelines {
		timelineRoleID := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		timelineID := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

		timelineMetadata := map[string]string{
			metadata.SubjectID:    usi,
			metadata.UserID:       usi,
			metadata.ResourceKind: "timeline",
			metadata.RoleKind:     "owner",
			metadata.RoleID:       timelineRoleID,
			metadata.TimelineID:   timelineID,
			metadata.VentureID:    ventureID,
		}

		{
			rolReq := role.CreateI{
				Obj: []*role.CreateI_Obj{
					{
						Metadata: timelineMetadata,
					},
				},
			}

			ok, err := h.storage.Role.Creator.Verify(ctx, &rolReq)
			if err != nil {
				return tracer.Mask(err)
			}

			if !ok {
				return tracer.Mask(invalidInputError)
			}

			_, err = h.storage.Role.Creator.Create(&rolReq)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			timelineReq := timeline.CreateI{
				Obj: []*timeline.CreateI_Obj{
					{
						Metadata: timelineMetadata,
						Property: &timeline.CreateI_Obj_Property{
							Desc: defaultTimeline.Desc,
							Name: defaultTimeline.Name,
						},
					},
				},
			}

			ok, err := h.storage.Timeline.Creator.Verify(ctx, &timelineReq)
			if err != nil {
				return tracer.Mask(err)
			}

			if !ok {
				return tracer.Mask(invalidInputError)
			}

			_, err = h.storage.Timeline.Creator.Create(&timelineReq)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
