package metric

import (
	"context"
	"fmt"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
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

func (h *Handler) Create(ctx context.Context, cre *metric.CreateI) (*metric.CreateO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.CreateO{}, nil
}

func (h *Handler) Delete(ctx context.Context, del *metric.DeleteI) (*metric.DeleteO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.DeleteO{}, nil
}

func (h *Handler) Search(ctx context.Context, sea *metric.SearchI) (*metric.SearchO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.SearchO{}, nil
}

func (h *Handler) Update(ctx context.Context, upd *metric.UpdateI) (*metric.UpdateO, error) {
	fmt.Printf("%#v\n", time.Now().String())
	return &metric.UpdateO{}, nil
}
