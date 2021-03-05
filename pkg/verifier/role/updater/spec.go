package updater

import "github.com/venturemark/apigengo/pkg/pbf/role"

type Interface interface {
	Verify(req *role.UpdateI) (bool, error)
}
