package audience

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/audience"
)

func (h *Handler) Delete(ctx context.Context, obj *audience.DeleteI) (*audience.DeleteO, error) {
	return &audience.DeleteO{}, nil
}
