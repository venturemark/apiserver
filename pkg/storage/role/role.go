package role

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/role/creator"
	"github.com/venturemark/apiserver/pkg/storage/role/deleter"
	"github.com/venturemark/apiserver/pkg/storage/role/searcher"
	"github.com/venturemark/apiserver/pkg/storage/role/updater"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
}

type Role struct {
	Creator  *creator.Creator
	Deleter  *deleter.Deleter
	Searcher *searcher.Searcher
	Updater  *updater.Updater
}

func New(config Config) (*Role, error) {
	var err error

	var cre *creator.Creator
	{
		c := creator.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
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
			Rescue: config.Rescue,
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
			Rescue: config.Rescue,
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
			Rescue: config.Rescue,
		}

		upd, err = updater.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	r := &Role{
		Creator:  cre,
		Deleter:  del,
		Searcher: sea,
		Updater:  upd,
	}

	return r, nil
}
