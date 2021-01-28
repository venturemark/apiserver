package metupd

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metupd/creator"
	"github.com/venturemark/apiserver/pkg/storage/metupd/updater"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
}

type MetUpd struct {
	Creator *creator.Creator
	Updater *updater.Updater
}

func New(config Config) (*MetUpd, error) {
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

	m := &MetUpd{
		Creator: cre,
		Updater: upd,
	}

	return m, nil
}
