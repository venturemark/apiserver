package deleter

import "github.com/venturemark/apigengo/pkg/pbf/texupd"

type Interface interface {
	Verify(req *texupd.DeleteI) (bool, error)
}