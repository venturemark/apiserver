package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/handler/metric/validate"
)

func (h *Handler) Update(ctx context.Context, obj *metric.UpdateI) (*metric.UpdateO, error) {
	var err error

	err = validate.Update(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	err = validate.Update(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return &metric.UpdateO{}, nil
}
