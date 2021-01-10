package audience

import (
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	audience.RegisterAPIServer(g, h)
}
