package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

func (h *Handler) Update(ctx context.Context, obj *timeline.UpdateI) (*timeline.UpdateO, error) {
	return &timeline.UpdateO{}, nil
}
