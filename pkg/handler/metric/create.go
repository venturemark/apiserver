package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/handler/metric/validate"
)

func (h *Handler) Create(ctx context.Context, obj *metric.CreateI) (*metric.CreateO, error) {
	var err error

	err = validate.Create(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	err = validate.Create(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return &metric.CreateO{}, nil
}
