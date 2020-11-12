package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/handler/metric/validate"
)

func (h *Handler) Delete(ctx context.Context, obj *metric.DeleteI) (*metric.DeleteO, error) {
	var err error

	err = validate.Delete(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	err = validate.Delete(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return &metric.DeleteO{}, nil
}
