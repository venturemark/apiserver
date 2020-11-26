package searcher

import "github.com/venturemark/apigengo/pkg/pbf/update"

type Interface interface {
	Verify(req *update.SearchI) (bool, error)
}
