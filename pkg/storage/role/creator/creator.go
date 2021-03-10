package creator

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/role/creator"
	"github.com/venturemark/apiserver/pkg/verifier/role/creator/empty"
)

type Config struct {
	Logger     logger.Interface
	Permission permission.Gateway
	Redigo     redigo.Interface
	Rescue     rescue.Interface
}

type Creator struct {
	logger logger.Interface
	redigo redigo.Interface
	rescue rescue.Interface

	verify []creator.Interface
}

func New(config Config) (*Creator, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Permission == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Permission must not be empty", config)
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

	c := &Creator{
		logger: config.Logger,
		redigo: config.Redigo,
		rescue: config.Rescue,

		verify: []creator.Interface{
			emptyVerifier,
		},
	}

	return c, nil
}
