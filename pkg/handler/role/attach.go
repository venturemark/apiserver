package role

import (
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	role.RegisterAPIServer(g, h)
}
