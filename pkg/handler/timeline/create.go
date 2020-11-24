package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

func (h *Handler) Create(ctx context.Context, obj *timeline.CreateI) (*timeline.CreateO, error) {
	return &timeline.CreateO{}, nil
}
