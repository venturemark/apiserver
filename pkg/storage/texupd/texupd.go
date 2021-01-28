package texupd

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/texupd/creator"
	"github.com/venturemark/apiserver/pkg/storage/texupd/deleter"
	"github.com/venturemark/apiserver/pkg/storage/texupd/updater"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
}

type TexUpd struct {
	Creator *creator.Creator
	Deleter *deleter.Deleter
	Updater *updater.Updater
}

func New(config Config) (*TexUpd, error) {
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

	t := &TexUpd{
		Creator: cre,
		Deleter: del,
		Updater: upd,
	}

	return t, nil
}
