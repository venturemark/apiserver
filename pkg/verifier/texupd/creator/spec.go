package creator

import "github.com/venturemark/apigengo/pkg/pbf/texupd"

type Interface interface {
	Verify(req *texupd.CreateI) (bool, error)
}
