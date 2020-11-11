package metric

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"

	"github.com/venturemark/apiserver/pkg/handler/metric/validate"
)

type HandlerConfig struct {
	Logger logger.Interface
}

type Handler struct {
	logger logger.Interface

	metric.UnimplementedAPIServer
}

func NewHandler(config HandlerConfig) (*Handler, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	h := &Handler{
		logger: config.Logger,
	}

	return h, nil
}

func (h *Handler) Attach(g *grpc.Server) {
	metric.RegisterAPIServer(g, h)
}

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

func (h *Handler) Search(ctx context.Context, obj *metric.SearchI) (*metric.SearchO, error) {
	var err error

	err = validate.Search(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	err = validate.Search(obj)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return &metric.SearchO{}, nil
}

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
