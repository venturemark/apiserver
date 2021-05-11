package deleter

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/role"
)

type Interface interface {
	Verify(ctx context.Context, req *role.DeleteI) (bool, error)
}
