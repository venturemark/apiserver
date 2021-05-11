package deleter

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

type Interface interface {
	Verify(ctx context.Context, req *timeline.DeleteI) (bool, error)
}
