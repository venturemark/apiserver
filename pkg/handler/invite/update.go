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
				req.Obj[i].Metadata[metadata.SubjectID] = usi
				req.Obj[i].Metadata[metadata.UserID] = usi
			}
		}
	}

	var res *invite.UpdateO
	{
		ok, err := h.storage.Invite.Updater.Verify(ctx, req)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if !ok {
			return nil, tracer.Mask(invalidInputError)
		}

		res, err = h.storage.Invite.Updater.Update(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		for i := range res.Obj {
			if res.Obj[i].Metadata[metadata.RoleStatus] == "created" {
				{
					req.Obj[i].Metadata[metadata.RoleID] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
				}

				{
					res.Obj[i].Metadata[metadata.RoleID] = req.Obj[i].Metadata[metadata.RoleID]
				}

				{
					rol := &role.CreateI{
						Obj: []*role.CreateI_Obj{
							{
								Metadata: req.Obj[i].Metadata,
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
			}
		}
	}

	return res, nil
}
