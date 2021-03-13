package updater

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/venture/updater"
	"github.com/venturemark/apiserver/pkg/verifier/venture/updater/auth"
	"github.com/venturemark/apiserver/pkg/verifier/venture/updater/patch"
)

type Config struct {
	Logger     logger.Interface
	Permission permission.Gateway
	Redigo     redigo.Interface
	Rescue     rescue.Interface
}

type Updater struct {
	logger logger.Interface
	redigo redigo.Interface
	rescue rescue.Interface

	verify []updater.Interface
}

func New(config Config) (*Updater, error) {
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

	var patchVerifier *patch.Verifier
	{
		c := patch.VerifierConfig{}

		patchVerifier, err = patch.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u := &Updater{
		logger: config.Logger,
		redigo: config.Redigo,
		rescue: config.Rescue,

		verify: []updater.Interface{
			authVerifier,
			patchVerifier,
		},
	}

	return u, nil
}
