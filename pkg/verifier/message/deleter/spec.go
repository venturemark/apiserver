package deleter

import "github.com/venturemark/apigengo/pkg/pbf/message"

type Interface interface {
	Verify(req *message.DeleteI) (bool, error)
}
