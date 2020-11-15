package update

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/update"
)

func (h *Handler) Create(ctx context.Context, obj *update.CreateI) (*update.CreateO, error) {
	return &update.CreateO{}, nil
}
