package searcher

import "github.com/venturemark/apigengo/pkg/pbf/invite"

type Interface interface {
	Verify(req *invite.SearchI) (bool, error)
}
