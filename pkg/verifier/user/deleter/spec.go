package deleter

import "github.com/venturemark/apigengo/pkg/pbf/user"

type Interface interface {
	Verify(req *user.DeleteI) (bool, error)
}
