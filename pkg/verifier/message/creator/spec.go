package creator

import "github.com/venturemark/apigengo/pkg/pbf/message"

type Interface interface {
	Verify(req *message.CreateI) (bool, error)
}
