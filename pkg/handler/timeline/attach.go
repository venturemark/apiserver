package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	timeline.RegisterAPIServer(g, h)
}
