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

	return true, nil
}
