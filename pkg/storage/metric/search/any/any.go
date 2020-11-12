package any

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metric/search/any/update"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Any struct {
	Update *update.Update
}

func New(config Config) (*Any, error) {
	var err error

	var u *update.Update
	{
		c := update.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		u, err = update.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	a := &Any{
		Update: u,
	}

	return a, nil
}
