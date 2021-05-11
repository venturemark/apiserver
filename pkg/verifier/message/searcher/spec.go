package searcher

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
)

type Interface interface {
	Verify(ctx context.Context, req *message.SearchI) (bool, error)
}
