package metupd

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

func (h *Handler) Search(ctx context.Context, obj *metupd.SearchI) (*metupd.SearchO, error) {
	return &metupd.SearchO{}, nil
}
