package timeline

import (
	"fmt"
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"
)

func (t *Timeline) Verify(obj *metupd.UpdateI) (bool, error) {
	{
		// Updating metrics is optional when updating metric updates. Somebody
		// may just wish to update their updates.
		if len(obj.Yaxis) != 0 {
			// We always check the latest item of the sorted set to check the amount
			// of datapoints on the y axis. Due to this very check the consistency
			// of the sorted set is ensured, which means that lookup up a single
			// element of the sorted set is sufficient.
			k := fmt.Sprintf("tml:%s:met", obj.Timeline)

			s, err := t.redigo.Scored().Search(k, 0, 1)
			if err != nil {
				return false, tracer.Mask(err)
			}

			// The elements of the sorted set are comma separated. Before comparing
			// the amount of y axis coordinates we need to remove the unix timestamp
			// from the list of elements, which is why we do -1 below. Then we
			// compare the amount of y axis coordinates given by the user input with
			// what we found in the sorted set. These two numbers must always match.
			// Otherwise the graphs on a timeline become incomprehensible.
			var c int
			var y int
			if len(s) == 1 {
				c = len(strings.Split(s[0], ",")) - 1
				y = len(obj.Yaxis)
			}

			if c != y {
				return false, nil
			}
		}
	}

	{
		// Updating updates is optional when updating metric updates. Somebody
		// may just wish to update their metrics. If the update text is
		// provided, it is still limited to up to 280 characters. Nobody should
		// be able to update metric updates with longer text.
		if obj.Text != "" && len(obj.Text) > 280 {
			return false, nil
		}
	}

	{
		// Updating metric updates requires either of the resources to be given.
		// It is not valid to request the update of any resource without
		// providing any of these resources.
		if len(obj.Yaxis) == 0 && obj.Text == "" {
			return false, nil
		}
	}

	{
		// Updating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if obj.Timeline == "" {
			return false, nil
		}
	}

	{
		// Updating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if obj.Timestamp == 0 {
			return false, nil
		}
	}

	return true, nil
}
