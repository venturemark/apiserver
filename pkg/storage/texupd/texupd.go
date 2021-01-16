package texupd

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/texupd/creator"
	"github.com/venturemark/apiserver/pkg/storage/texupd/updater"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type TexUpd struct {
	Creator *creator.Creator
	Updater *updater.Updater
}

func New(config Config) (*TexUpd, error) {
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

	t := &TexUpd{
		Creator: cre,
		Updater: upd,
	}

	return t, nil
}