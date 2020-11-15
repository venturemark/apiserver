package update

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/update/search"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Update struct {
	Search *search.Search
}

func New(config Config) (*Update, error) {
	var err error

	var s *search.Search
	{
		c := search.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		s, err = search.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u := &Update{
		Search: s,
	}

	return u, nil
}
