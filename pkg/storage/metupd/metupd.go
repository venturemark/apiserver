package metupd

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metupd/create"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Metric struct {
	Create *create.Create
}

func New(config Config) (*Metric, error) {
	var err error

	var c *create.Create
	{
		c := create.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		c, err = create.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	m := &Metric{
		Create: c,
	}

	return m, nil
}
