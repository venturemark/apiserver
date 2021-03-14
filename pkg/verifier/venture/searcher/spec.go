package searcher

import "github.com/venturemark/apigengo/pkg/pbf/venture"

type Interface interface {
	Verify(req *venture.SearchI) (bool, error)
}
