package searcher

import "github.com/venturemark/apigengo/pkg/pbf/role"

type Interface interface {
	Verify(req *role.SearchI) (bool, error)
}
