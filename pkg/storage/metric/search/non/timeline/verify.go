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
		// Search requests without operator should be prohibited. The client
		// needs to tell us exactly what we need to do.
		if len(obj.Filter.Operator) != 1 {
			return false
		}

		// This particular search implementation is to search for any matching
		// property we can find. Anything else but the any operator is
		// prohibited.
		if obj.Filter.Operator[0] != "any" {
			return false
		}
	}

	{
		// Searching without having any property given we can use to search for
		// does not work. We always need to receive at least one property.
		if len(obj.Filter.Property) < 2 {
			return false
		}

		for _, p := range obj.Filter.Property {
			// It is not allowed to provide timestamp properties with the search
			// request of this particular search implementation.
			if p.Timestamp != "" {
				return false
			}
			// With this particular search implementation we require only
			// timeline IDs to be given. If any timeline ID property is empty,
			// we decline service for this request.
			if p.Timeline == "" {
				return false
			}
		}
	}

	{
		var l []string

		for _, p := range obj.Filter.Property {
			l = append(l, p.Timeline)
		}

		// We do not want to do unnecessary work. We want clients to be aware of
		// the search requests they are sending. Therefore we deny any
		// duplicated properties.
		if !unique(l) {
			return false
		}
	}

	return true
}

func unique(l []string) bool {
	for _, s := range l {
		if count(l, s) > 1 {
			return false
		}
	}

	return true
}

func count(l []string, s string) int {
	var c int

	for _, i := range l {
		if i == s {
			c++
		}
	}

	return c
}
