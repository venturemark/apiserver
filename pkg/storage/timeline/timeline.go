package timeline

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/timeline/search"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Timeline struct {
	Search *search.Search
}

func New(config Config) (*Timeline, error) {
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

	t := &Timeline{
		Search: s,
	}

	return t, nil
}
