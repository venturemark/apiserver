package create

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type MetUpd struct {
	logger logger.Interface
	regigo redigo.Interface
}

func New(config Config) (*MetUpd, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	m := &MetUpd{
		logger: config.Logger,
		regigo: config.Redigo,
	}

	return m, nil
}
