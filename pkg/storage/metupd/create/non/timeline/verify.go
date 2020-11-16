package timeline

import (
	"fmt"
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/tracer"
)

func (t *Timeline) Verify(obj *metupd.CreateI) (bool, error) {
	{
		// Creating metric updates requires at least one coordinate on the y
		// axis. If the client does not provide that we fail.
		if len(obj.Yaxis) == 0 {
			return false, nil
		}

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

	{
		// Creating metric updates requires some text to be provided. Without
		// text we cannot allow a metric update to be created.
		if obj.Text == "" {
			return false, nil
		}

		// Creating metric updates is limited to text with up to 280 characters.
		// Nobody should be able to create metric updates with longer text.
		if len(obj.Text) > 280 {
			return false, nil
		}
	}

	{
		// Creating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if obj.Timeline == "" {
			return false, nil
		}
	}

	return true, nil
}
