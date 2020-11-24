package update

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/update/searcher"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Update struct {
	Searcher *searcher.Searcher
}

func New(config Config) (*Update, error) {
	var err error

	var s *searcher.Searcher
	{
		c := searcher.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		s, err = searcher.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u := &Update{
		Searcher: s,
	}

	return u, nil
}
