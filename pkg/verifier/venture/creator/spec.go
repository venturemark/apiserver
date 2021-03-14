package creator

import "github.com/venturemark/apigengo/pkg/pbf/venture"

type Interface interface {
	Verify(req *venture.CreateI) (bool, error)
}
