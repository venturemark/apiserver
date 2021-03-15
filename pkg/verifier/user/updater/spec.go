package updater

import "github.com/venturemark/apigengo/pkg/pbf/user"

type Interface interface {
	Verify(req *user.UpdateI) (bool, error)
}
