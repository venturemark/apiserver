package searcher

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/role"
)

type Interface interface {
	Verify(ctx context.Context, req *role.SearchI) (bool, error)
}
