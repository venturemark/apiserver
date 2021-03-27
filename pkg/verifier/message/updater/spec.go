package updater

import "github.com/venturemark/apigengo/pkg/pbf/message"

type Interface interface {
	Verify(req *message.UpdateI) (bool, error)
}
