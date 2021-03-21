package creator

import "github.com/venturemark/apigengo/pkg/pbf/invite"

type Interface interface {
	Verify(req *invite.CreateI) (bool, error)
}
