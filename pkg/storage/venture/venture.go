package venture

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/venture/creator"
	"github.com/venturemark/apiserver/pkg/storage/venture/deleter"
	"github.com/venturemark/apiserver/pkg/storage/venture/searcher"
	"github.com/venturemark/apiserver/pkg/storage/venture/updater"
)

type Config struct {
	Logger     logger.Interface
	Permission permission.Gateway
	Redigo     redigo.Interface
	Rescue     rescue.Interface
}

type Venture struct {
	Creator  *creator.Creator
	Deleter  *deleter.Deleter
	Searcher *searcher.Searcher
	Updater  *updater.Updater
}

func New(config Config) (*Venture, error) {
	var err error

	var cre *creator.Creator
	{
		c := creator.Config{
			Logger:     config.Logger,
			Permission: config.Permission,
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
			Logger:     config.Logger,
			Permission: config.Permission,
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
			Logger:     config.Logger,
			Permission: config.Permission,
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
			Logger:     config.Logger,
			Permission: config.Permission,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		upd, err = updater.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	v := &Venture{
		Creator:  cre,
		Deleter:  del,
		Searcher: sea,
		Updater:  upd,
	}

	return v, nil
}
