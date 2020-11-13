package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
)

func (h *Handler) Delete(ctx context.Context, obj *metric.DeleteI) (*metric.DeleteO, error) {
	return &metric.DeleteO{}, nil
}
