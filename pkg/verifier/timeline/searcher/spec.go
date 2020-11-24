package searcher

import "github.com/venturemark/apigengo/pkg/pbf/timeline"

type Interface interface {
	Verify(req *timeline.SearchI) (bool, error)
}
