package message

import (
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	message.RegisterAPIServer(g, h)
}
