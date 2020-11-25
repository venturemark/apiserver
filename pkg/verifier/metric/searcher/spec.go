package searcher

import "github.com/venturemark/apigengo/pkg/pbf/metric"

type Interface interface {
	Verify(req *metric.SearchI) (bool, error)
}
