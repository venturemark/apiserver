package metupd

import (
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	metupd.RegisterAPIServer(g, h)
}
