package metupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

func (h *Handler) Update(ctx context.Context, obj *metupd.UpdateI) (*metupd.UpdateO, error) {
	return &metupd.UpdateO{}, nil
}
