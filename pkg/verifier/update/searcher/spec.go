package searcher

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/update"
)

type Interface interface {
	Verify(ctx context.Context, req *update.SearchI) (bool, error)
}
