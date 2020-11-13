package storage

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metric"
	"github.com/venturemark/apiserver/pkg/storage/metupd"
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

	var metricStorage *metric.Metric
	{
		c := metric.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		metricStorage, err = metric.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var metupdStorage *metupd.Metric
	{
		c := metupd.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		metupdStorage, err = metupd.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	s := &Storage{
		Metric: metricStorage,
		MetUpd: metupdStorage,
	}

	return s, nil
}
