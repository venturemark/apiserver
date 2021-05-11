package updater

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/invite"
)

type Interface interface {
	Verify(ctx context.Context, req *invite.UpdateI) (bool, error)
}
