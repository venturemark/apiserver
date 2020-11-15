package update

import (
	"github.com/venturemark/apigengo/pkg/pbf/update"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	update.RegisterAPIServer(g, h)
}
