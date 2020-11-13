package non

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metric/search/non/timeline"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Non struct {
	Timeline *timeline.Timeline
}

func New(config Config) (*Non, error) {
	var err error

	var t *timeline.Timeline
	{
		c := timeline.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		t, err = timeline.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	n := &Non{
		Timeline: t,
	}

	return n, nil
}
