package storage

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metric"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Storage struct {
	Metric *metric.Metric
}

func New(config Config) (*Storage, error) {
	var err error

	var m *metric.Metric
	{
		c := metric.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		m, err = metric.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	s := &Storage{
		Metric: m,
	}

	return s, nil
}
