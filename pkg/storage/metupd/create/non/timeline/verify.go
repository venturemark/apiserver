package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

func (t *Timeline) Verify(obj *metupd.CreateI) bool {
	{
		// Creating metric updates requires two datapoints. If the client
		// does not provide that we fail.
		if len(obj.Datapoint) != 2 {
			return false
		}
	}

	{
		// Creating metric updates requires some text to be provided. Without
		// text we cannot allow a metric update to be created.
		if obj.Text == "" {
			return false
		}

		// Creating metric updates is limited to text with up to 280 characters.
		// Nobody should be able to create metric updates with longer text.
		if len(obj.Text) > 280 {
			return false
		}
	}

	{ // nolint: gosimple
		// Creating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if obj.Timeline == "" {
			return false
		}
	}

	return true
}
