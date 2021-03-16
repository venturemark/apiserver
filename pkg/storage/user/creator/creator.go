package creator

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/association"
	"github.com/venturemark/apiserver/pkg/verifier/user/creator"
	"github.com/venturemark/apiserver/pkg/verifier/user/creator/auth"
	"github.com/venturemark/apiserver/pkg/verifier/user/creator/empty"
	"github.com/venturemark/apiserver/pkg/verifier/user/creator/name"
)

type Config struct {
	Association *association.Association
	Logger      logger.Interface
	Permission  permission.Gateway
	Redigo      redigo.Interface
	Rescue      rescue.Interface
}

type Creator struct {
	association *association.Association
	logger      logger.Interface
	redigo      redigo.Interface
	rescue      rescue.Interface

	verify []creator.Interface
}

func New(config Config) (*Creator, error) {
	if config.Association == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Association must not be empty", config)
	}
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

	var nameVerifier *name.Verifier
	{
		c := name.VerifierConfig{}

		nameVerifier, err = name.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	c := &Creator{
		association: config.Association,
		logger:      config.Logger,
		redigo:      config.Redigo,
		rescue:      config.Rescue,

		verify: []creator.Interface{
			// The empty verifier must be run first so that following verifiers
			// do not have to check for prerequisites over and over again.
			emptyVerifier,
			authVerifier,
			nameVerifier,
		},
	}

	return c, nil
}
