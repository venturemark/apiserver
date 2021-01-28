package deleter

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/message/deleter"
	"github.com/venturemark/apiserver/pkg/verifier/message/deleter/empty"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
}

type Deleter struct {
	logger logger.Interface
	redigo redigo.Interface
	rescue rescue.Interface

	verify []deleter.Interface
}

func New(config Config) (*Deleter, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}
	if config.Rescue == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Rescue must not be empty", config)
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
		rescue: config.Rescue,

		verify: []deleter.Interface{
			emptyVerifier,
		},
	}

	return d, nil
}
