package texupd

import (
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	texupd.RegisterAPIServer(g, h)
}
