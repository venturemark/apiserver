package timeline

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/metric/search"
	"github.com/venturemark/apiserver/pkg/verifier/metric/search/empty"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Timeline struct {
	logger logger.Interface
	redigo redigo.Interface

	verify []search.Interface
}

func New(config Config) (*Timeline, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	var err error

	var emptyVerifier *empty.Verifier
	{
		c := empty.VerifierConfig{}

		emptyVerifier, err = empty.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	t := &Timeline{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []search.Interface{
			emptyVerifier,
		},
	}

	return t, nil
}
