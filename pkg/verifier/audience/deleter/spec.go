package deleter

import "github.com/venturemark/apigengo/pkg/pbf/audience"

type Interface interface {
	Verify(req *audience.DeleteI) (bool, error)
}
