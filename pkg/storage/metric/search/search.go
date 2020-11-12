package search

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metric/search/any"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Search struct {
	Any *any.Any
}

func New(config Config) (*Search, error) {
	var err error

	var a *any.Any
	{
		c := any.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		a, err = any.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	s := &Search{
		Any: a,
	}

	return s, nil
}
