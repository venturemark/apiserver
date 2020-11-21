package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func (t *Timeline) Verify(req *metupd.UpdateI) (bool, error) {
	for _, v := range t.verify {
		ok, err := v.Verify(req)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if !ok {
			return false, nil
		}
	}

	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}

		// Updating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if req.Obj.Metadata[metadata.Timeline] == "" {
			return false, nil
		}

		// Updating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if req.Obj.Metadata[metadata.Unixtime] == "" {
			return false, nil
		}
	}

	{
		// Updating updates is optional when updating metric updates. Somebody
		// may just wish to update their metrics. If the update text is
		// provided, it is still limited to up to 280 characters. Nobody should
		// be able to update metric updates with longer text.
		if req.Obj.Property.Text != "" && len(req.Obj.Property.Text) > 280 {
			return false, nil
		}
	}

	{
		// Updating metrics is optional when updating metric updates. Somebody
		// may just wish to update their updates.
		if len(req.Obj.Property.Data) != 0 {
			// We do this step separately for reasons of performance and impact on
			// the operational system. We do not need to execute any further checks
			// if the provided datastructure is already insufficient.
			for _, d := range req.Obj.Property.Data {
				if len(d.Value) == 0 {
					return false, nil
				}
			}

			// Dimensional spaces must be identified with single character
			// variables. Anything else other than x, y, z is invalid. Additionally
			// the reserved dimensional space t must also not be supplied since the
			// system provides that automatically.
			for _, d := range req.Obj.Property.Data {
				if len(d.Space) != 1 {
					return false, nil
				}
				if d.Space == "t" {
					return false, nil
				}
			}

			// We do not permit updating datapoints for the same dimensional
			// space twice. If the user tries to update a metric update with
			// e.g. the dimension y not being unique, the request fails.
			for i, d := range req.Obj.Property.Data {
				if i == 0 {
					continue
				}
				if req.Obj.Property.Data[0].Space == d.Space {
					return false, nil
				}
			}

			// The amount of all datapoints must be equal across dimensions
			// provided. We do not permit inconsistencies within the request data.
			for i, d := range req.Obj.Property.Data {
				if i == 0 {
					continue
				}
				if len(req.Obj.Property.Data[0].Value) != len(d.Value) {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
