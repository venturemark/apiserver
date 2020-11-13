package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
)

func (t *Timeline) Verify(obj *metric.SearchI) bool {
	{
		// We need a filter. Any search request without it does not make sense,
		// because we do then not even know what we should search for.
		if obj.Filter == nil {
			return false
		}
	}

	{
		// Chunking is not supported at this point. Any usage of something that
		// is not implemented should be prohibited in order to not create the
		// wrong expectations.
		if obj.Filter.Chunking != nil {
			return false
		}
	}

	{
		// Search requests must not provide any operator for this
		// implementation. The client tells us exactly what we need to do with
		// the single timeline ID they provide.
		if len(obj.Filter.Operator) != 0 {
			return false
		}
	}

	{
		// Searching using this implementation requires only a single timeline
		// ID to be given.
		if len(obj.Filter.Property) != 1 {
			return false
		}

		for _, p := range obj.Filter.Property {
			// It is not allowed to provide timestamp properties with the search
			// request of this particular search implementation.
			if p.Timestamp != 0 {
				return false
			}
			// With this particular search implementation we require only a
			// single timeline ID to be given. If the timeline ID property is
			// empty, we decline service for this request.
			if p.Timeline == "" {
				return false
			}
		}
	}

	return true
}
