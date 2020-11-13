package metupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

func (h *Handler) Delete(ctx context.Context, obj *metupd.DeleteI) (*metupd.DeleteO, error) {
	return &metupd.DeleteO{}, nil
}
