package timeline

import "github.com/venturemark/apigengo/pkg/pbf/metupd"

type Verifier interface {
	Verify(req *metupd.UpdateI) (bool, error)
}
