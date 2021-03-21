package invite

import (
	"github.com/venturemark/apigengo/pkg/pbf/invite"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	invite.RegisterAPIServer(g, h)
}
