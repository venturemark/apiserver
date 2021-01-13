package searcher

import "github.com/venturemark/apigengo/pkg/pbf/message"

type Interface interface {
	Verify(req *message.SearchI) (bool, error)
}
