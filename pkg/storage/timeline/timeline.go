package timeline

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/timeline/creator"
	"github.com/venturemark/apiserver/pkg/storage/timeline/searcher"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Timeline struct {
	Creator  *creator.Creator
	Searcher *searcher.Searcher
}

func New(config Config) (*Timeline, error) {
	var err error

	var cre *creator.Creator
	{
		c := creator.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		cre, err = creator.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sea *searcher.Searcher
	{
		c := searcher.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		sea, err = searcher.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	t := &Timeline{
		Creator:  cre,
		Searcher: sea,
	}

	return t, nil
}
