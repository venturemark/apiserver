package venture

import (
	"github.com/venturemark/apigengo/pkg/pbf/venture"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	venture.RegisterAPIServer(g, h)
}
