package metric

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	metric.RegisterAPIServer(g, h)
}
