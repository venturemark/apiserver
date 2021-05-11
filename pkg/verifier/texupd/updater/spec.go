package updater

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
)

type Interface interface {
	Verify(ctx context.Context, req *texupd.UpdateI) (bool, error)
}
