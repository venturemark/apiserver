package audience

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/audience/creator"
	"github.com/venturemark/apiserver/pkg/storage/audience/deleter"
	"github.com/venturemark/apiserver/pkg/storage/audience/searcher"
	"github.com/venturemark/apiserver/pkg/storage/audience/updater"
)

type Config struct {
	Permission permission.Gateway
	Logger     logger.Interface
	Redigo     redigo.Interface
	Rescue     rescue.Interface
}

type Audience struct {
	Creator  *creator.Creator
	Deleter  *deleter.Deleter
	Searcher *searcher.Searcher
	Updater  *updater.Updater
}

func New(config Config) (*Audience, error) {
	var err error

	var cre *creator.Creator
	{
		c := creator.Config{
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		cre, err = creator.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var del *deleter.Deleter
	{
		c := deleter.Config{
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		del, err = deleter.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sea *searcher.Searcher
	{
		c := searcher.Config{
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		sea, err = searcher.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var upd *updater.Updater
	{
		c := updater.Config{
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		upd, err = updater.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	a := &Audience{
		Creator:  cre,
		Deleter:  del,
		Searcher: sea,
		Updater:  upd,
	}

	return a, nil
}
