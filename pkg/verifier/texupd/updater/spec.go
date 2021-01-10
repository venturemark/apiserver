package updater

import "github.com/venturemark/apigengo/pkg/pbf/texupd"

type Interface interface {
	Verify(req *texupd.UpdateI) (bool, error)
}
