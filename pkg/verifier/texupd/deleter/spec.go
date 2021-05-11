package deleter

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
)

type Interface interface {
	Verify(ctx context.Context, req *texupd.DeleteI) (bool, error)
}
