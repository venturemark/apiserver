package timeline

import (
	"fmt"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/metric/timeline/data"
)

func (t *Timeline) Verify(req *metupd.CreateI) (bool, error) {
	{
		// Creating metric updates requires a timeline ID to be provided. This
		// is used as common denominator for the metric and the update resource
		// respectively. If the timeline ID is not given with the object
		// metadata, we decline service for this request.
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
		if req.Obj.Metadata[metadata.Timeline] == "" {
			return false, nil
		}
	}

	{
		// Creating metric updates requires at least one datapoint on any
		// dimension. If the client does not provide that we fail.
		if len(req.Obj.Property.Data) == 0 {
			return false, nil
		}

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

		// We do not permit creating datapoints for the same dimensional space
		// twice. If the user tries to create a metric update with e.g. the
		// dimension y not being unique, the request fails.
		for i, d := range req.Obj.Property.Data {
			if i == 0 {
				continue
			}
			if req.Obj.Property.Data[0].Space == d.Space {
				return false, nil
			}
		}

		// The amount of all datapoints must be equal across dimensions
		// provided. We do not permit inconsistencies with the request data.
		for i, d := range req.Obj.Property.Data {
			if i == 0 {
				continue
			}
			if len(req.Obj.Property.Data[0].Value) != len(d.Value) {
				return false, nil
			}
		}

		// We always check the amount of datapoints on the first dimension given
		// compared to the data already stored. Due to this very check the
		// consistency of the sorted set is ensured. The amount of datapoints
		// tracked per dimension must always match. Otherwise the graphs on a
		// timeline become incomprehensible. Note that the consistency of the
		// dimensions given is already verified above. This means we can resort
		// to only verify the first dimension given.
		k := fmt.Sprintf(key.Timeline, req.Obj.Metadata[metadata.Timeline])
		s, err := t.redigo.Scored().Search(k, 0, 1)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if len(s) == 1 {
			_, v, err := data.Split(s[0])
			if err != nil {
				return false, tracer.Mask(err)
			}

			c := len(v[0].GetValue())
			y := len(req.Obj.Property.Data[0].Value)
			if c != y {
				return false, nil
			}
		}
	}

	{
		// Creating metric updates requires some text to be provided. Without
		// text we cannot allow a metric update to be created.
		if req.Obj.Property.Text == "" {
			return false, nil
		}

		// Creating metric updates is limited to text with up to 280 characters.
		// Nobody should be able to create metric updates with longer text.
		if len(req.Obj.Property.Text) > 280 {
			return false, nil
		}
	}

	return true, nil
}
