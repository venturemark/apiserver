package audience

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/audience/creator"
	"github.com/venturemark/apiserver/pkg/storage/audience/searcher"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Audience struct {
	Creator  *creator.Creator
	Searcher *searcher.Searcher
}

func New(config Config) (*Audience, error) {
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

	a := &Audience{
		Creator:  cre,
		Searcher: sea,
	}

	return a, nil
}
