package audience

import (
	"github.com/venturemark/apigengo/pkg/pbf/audience"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage"
)

type HandlerConfig struct {
	Logger  logger.Interface
	Storage *storage.Storage
}

type Handler struct {
	logger  logger.Interface
	storage *storage.Storage

	audience.UnimplementedAPIServer
}

func NewHandler(config HandlerConfig) (*Handler, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Storage == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Storage must not be empty", config)
	}

	h := &Handler{
		logger:  config.Logger,
		storage: config.Storage,
	}

	return h, nil
}
