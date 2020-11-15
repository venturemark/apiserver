package update

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/update"
)

func (h *Handler) Delete(ctx context.Context, obj *update.DeleteI) (*update.DeleteO, error) {
	return &update.DeleteO{}, nil
}
