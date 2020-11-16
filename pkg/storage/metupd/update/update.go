package update

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metupd/update/non"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Update struct {
	Non *non.Non
}

func New(config Config) (*Update, error) {
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

	u := &Update{
		Non: n,
	}

	return u, nil
}
