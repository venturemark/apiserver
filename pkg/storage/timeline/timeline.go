package timeline

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/timeline/creator"
	"github.com/venturemark/apiserver/pkg/storage/timeline/deleter"
	"github.com/venturemark/apiserver/pkg/storage/timeline/searcher"
	"github.com/venturemark/apiserver/pkg/storage/timeline/updater"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Timeline struct {
	Creator  *creator.Creator
	Deleter  *deleter.Deleter
	Searcher *searcher.Searcher
	Updater  *updater.Updater
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

	var del *deleter.Deleter
	{
		c := deleter.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		del, err = deleter.New(c)
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

	var upd *updater.Updater
	{
		c := updater.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		upd, err = updater.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	t := &Timeline{
		Creator:  cre,
		Deleter:  del,
		Searcher: sea,
		Updater:  upd,
	}

	return t, nil
}
