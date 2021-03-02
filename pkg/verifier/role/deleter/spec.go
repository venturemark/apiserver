package deleter

import "github.com/venturemark/apigengo/pkg/pbf/role"

type Interface interface {
	Verify(req *role.DeleteI) (bool, error)
}
