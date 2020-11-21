package timeline

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/metupd/update/non/timeline/verify/consistency"
	"github.com/venturemark/apiserver/pkg/storage/metupd/update/non/timeline/verify/time"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Timeline struct {
	logger logger.Interface
	redigo redigo.Interface

	verify []Verifier
}

func New(config Config) (*Timeline, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	var err error

	var consistencyVerifier *consistency.Verifier
	{
		c := consistency.VerifierConfig{
			Redigo: config.Redigo,
		}

		consistencyVerifier, err = consistency.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var timeVerifier *time.Verifier
	{
		c := time.VerifierConfig{
			Now: now(),
		}

		timeVerifier, err = time.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	t := &Timeline{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []Verifier{
			consistencyVerifier,
			timeVerifier,
		},
	}

	return t, nil
}
