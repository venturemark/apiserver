package create

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metupd/create/non"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Create struct {
	Non *non.Non
}

func New(config Config) (*Create, error) {
	var err error

	var n *non.Non
	{
		c := non.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		n, err = non.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	s := &Create{
		Non: n,
	}

	return s, nil
}
