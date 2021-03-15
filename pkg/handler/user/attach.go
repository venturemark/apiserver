package user

import (
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	user.RegisterAPIServer(g, h)
}
