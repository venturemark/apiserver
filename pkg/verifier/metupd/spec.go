package metupd

import "github.com/venturemark/apigengo/pkg/pbf/metupd"

type Interface interface {
	Verify(req *metupd.UpdateI) (bool, error)
}
