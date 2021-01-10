package creator

import "github.com/venturemark/apigengo/pkg/pbf/audience"

type Interface interface {
	Verify(req *audience.CreateI) (bool, error)
}
