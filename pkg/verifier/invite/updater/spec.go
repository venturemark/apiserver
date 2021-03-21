package updater

import "github.com/venturemark/apigengo/pkg/pbf/invite"

type Interface interface {
	Verify(req *invite.UpdateI) (bool, error)
}
