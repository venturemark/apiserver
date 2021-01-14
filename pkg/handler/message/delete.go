package message

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/message"
)

func (h *Handler) Delete(ctx context.Context, obj *message.DeleteI) (*message.DeleteO, error) {
	return &message.DeleteO{}, nil
}
