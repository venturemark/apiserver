package searcher

import "github.com/venturemark/apigengo/pkg/pbf/user"

type Interface interface {
	Verify(req *user.SearchI) (bool, error)
}
