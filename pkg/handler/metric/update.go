package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
)

func (h *Handler) Update(ctx context.Context, obj *metric.UpdateI) (*metric.UpdateO, error) {
	return &metric.UpdateO{}, nil
}
