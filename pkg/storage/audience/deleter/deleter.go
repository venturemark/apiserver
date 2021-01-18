package deleter

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/audience/deleter"
	"github.com/venturemark/apiserver/pkg/verifier/audience/deleter/empty"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Deleter struct {
	logger logger.Interface
	redigo redigo.Interface

	verify []deleter.Interface
}

func New(config Config) (*Deleter, error) {
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

	d := &Deleter{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []deleter.Interface{
			emptyVerifier,
		},
	}

	return d, nil
}
