package deleter

import "github.com/venturemark/apigengo/pkg/pbf/invite"

type Interface interface {
	Verify(req *invite.DeleteI) (bool, error)
}
