package updater

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/role"
)

type Interface interface {
	Verify(ctx context.Context, req *role.UpdateI) (bool, error)
}
