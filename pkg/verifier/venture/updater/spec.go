package updater

import "github.com/venturemark/apigengo/pkg/pbf/venture"

type Interface interface {
	Verify(req *venture.UpdateI) (bool, error)
}
