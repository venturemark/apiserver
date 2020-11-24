package metupd

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metupd/creator"
	"github.com/venturemark/apiserver/pkg/storage/metupd/update"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type MetUpd struct {
	Creator *creator.Creator
	Update  *update.Update
}

func New(config Config) (*MetUpd, error) {
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

	var upd *update.Update
	{
		c := update.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		upd, err = update.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	m := &MetUpd{
		Creator: cre,
		Update:  upd,
	}

	return m, nil
}
