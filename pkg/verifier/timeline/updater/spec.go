package updater

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

type Interface interface {
	Verify(ctx context.Context, req *timeline.UpdateI) (bool, error)
}
