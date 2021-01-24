package updater

import "github.com/venturemark/apigengo/pkg/pbf/audience"

type Interface interface {
	Verify(req *audience.UpdateI) (bool, error)
}
