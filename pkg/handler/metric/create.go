package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
)

func (h *Handler) Create(ctx context.Context, obj *metric.CreateI) (*metric.CreateO, error) {
	return &metric.CreateO{}, nil
}
