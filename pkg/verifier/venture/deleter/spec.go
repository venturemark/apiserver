package deleter

import "github.com/venturemark/apigengo/pkg/pbf/venture"

type Interface interface {
	Verify(req *venture.DeleteI) (bool, error)
}
