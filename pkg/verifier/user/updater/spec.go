package updater

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/user"
)

type Interface interface {
	Verify(ctx context.Context, req *user.UpdateI) (bool, error)
}
