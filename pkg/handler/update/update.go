package update

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/update"
)

func (h *Handler) Update(ctx context.Context, req *update.UpdateI) (*update.UpdateO, error) {
	return &update.UpdateO{}, nil
}
