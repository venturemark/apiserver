package timeline

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Timeline struct {
	logger logger.Interface
	redigo redigo.Interface
}

func New(config Config) (*Timeline, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	t := &Timeline{
		logger: config.Logger,
		redigo: config.Redigo,
	}

	return t, nil
}
