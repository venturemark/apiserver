package creator

import "github.com/venturemark/apigengo/pkg/pbf/user"

type Interface interface {
	Verify(req *user.CreateI) (bool, error)
}
