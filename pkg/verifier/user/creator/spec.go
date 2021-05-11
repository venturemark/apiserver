package creator

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/user"
)

type Interface interface {
	Verify(ctx context.Context, req *user.CreateI) (bool, error)
}
