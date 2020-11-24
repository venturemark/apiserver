package timeline

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

func (h *Handler) Delete(ctx context.Context, obj *timeline.DeleteI) (*timeline.DeleteO, error) {
	return &timeline.DeleteO{}, nil
}
