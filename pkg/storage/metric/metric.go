package metric

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metric/searcher"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
}

type Metric struct {
	Searcher *searcher.Searcher
}

func New(config Config) (*Metric, error) {
	var err error

	var s *searcher.Searcher
	{
		c := searcher.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
			Rescue: config.Rescue,
		}

		s, err = searcher.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	m := &Metric{
		Searcher: s,
	}

	return m, nil
}
