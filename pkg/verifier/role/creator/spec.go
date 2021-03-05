package creator

import "github.com/venturemark/apigengo/pkg/pbf/role"

type Interface interface {
	Verify(req *role.CreateI) (bool, error)
}
