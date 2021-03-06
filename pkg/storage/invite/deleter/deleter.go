package deleter

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/invite/deleter"
	"github.com/venturemark/apiserver/pkg/verifier/invite/deleter/auth"
	"github.com/venturemark/apiserver/pkg/verifier/invite/deleter/empty"
	"github.com/venturemark/apiserver/pkg/verifier/invite/deleter/exist"
)

type Config struct {
	Logger     logger.Interface
	Permission permission.Gateway
	Redigo     redigo.Interface
	Rescue     rescue.Interface
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

	var authVerifier *auth.Verifier
	{
		c := auth.VerifierConfig{
			Permission: config.Permission,
			Redigo:     config.Redigo,
		}

		authVerifier, err = auth.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var emptyVerifier *empty.Verifier
	{
		c := empty.VerifierConfig{}

		emptyVerifier, err = empty.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var existVerifier *exist.Verifier
	{
		c := exist.VerifierConfig{
			Redigo: config.Redigo,
		}

		existVerifier, err = exist.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	d := &Deleter{
		logger: config.Logger,
		redigo: config.Redigo,
		rescue: config.Rescue,

		verify: []deleter.Interface{
			// The empty verifier must be run first so that following verifiers
			// do not have to check for prerequisites over and over again.
			emptyVerifier,
			authVerifier,
			existVerifier,
		},
	}

	return d, nil
}
