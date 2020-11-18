package timeline

import (
	"fmt"
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
)

func (t *Timeline) Verify(req *metupd.CreateI) (bool, error) {
	{
		// Creating metric updates requires a timeline ID to be provided. This
		// is used as common denominator for the metric and the update resource
		// respectively. If the timeline ID is not given with the object
		// metadata, we decline service for this request.
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

		// We always check the amount of datapoints on any dimension given. Due
		// to this very check the consistency of the sorted set is ensured,
		// which means that lookup up a single element of the sorted set is
		// sufficient.
		k := fmt.Sprintf(key.Timeline, req.Obj.Metadata[metadata.Timeline])
		s, err := t.redigo.Scored().Search(k, 0, 1)
		if err != nil {
			return false, tracer.Mask(err)
		}

		// TODO

		// The elements of the sorted set are comma separated. Before
		// comparing the amount of datapoints on each dimensional space we
		// need to remove the unix timestamp from the element, which is why
		// we do -1 on the string split below. Then we compare the amount of
		// values given by the user input with what we found in the sorted
		// set. These two numbers must always match. Otherwise the graphs on
		// a timeline become incomprehensible.
		for _, d := range req.Obj.Property.Data {
			var c int
			var y int
			if len(s) == 1 {
				c = len(strings.Split(s[0], ",")) - 1
				y = len(d.Value)
			}

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
