package creator

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/venture"
)

type Interface interface {
	Verify(ctx context.Context, req *venture.CreateI) (bool, error)
}
