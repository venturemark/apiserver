package updater

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
)

type Interface interface {
	Verify(ctx context.Context, req *message.UpdateI) (bool, error)
}
