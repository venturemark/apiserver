package message

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
)

func (h *Handler) Update(ctx context.Context, obj *message.UpdateI) (*message.UpdateO, error) {
	return &message.UpdateO{}, nil
}
