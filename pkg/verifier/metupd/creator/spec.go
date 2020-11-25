package creator

import "github.com/venturemark/apigengo/pkg/pbf/metupd"

type Interface interface {
	Verify(req *metupd.CreateI) (bool, error)
}
