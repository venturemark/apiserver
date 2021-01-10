package searcher

import "github.com/venturemark/apigengo/pkg/pbf/audience"

type Interface interface {
	Verify(req *audience.SearchI) (bool, error)
}
