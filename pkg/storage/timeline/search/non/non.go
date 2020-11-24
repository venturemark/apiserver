package non

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/timeline/search/non/user"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Non struct {
	User *user.User
}

func New(config Config) (*Non, error) {
	var err error

	var u *user.User
	{
		c := user.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		u, err = user.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	n := &Non{
		User: u,
	}

	return n, nil
}
