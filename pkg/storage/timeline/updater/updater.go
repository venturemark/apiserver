package updater

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/timeline/updater"
	"github.com/venturemark/apiserver/pkg/verifier/timeline/updater/empty"
	"github.com/venturemark/apiserver/pkg/verifier/timeline/updater/length"
	"github.com/venturemark/apiserver/pkg/verifier/timeline/updater/state"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
	Rescue rescue.Interface
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

	var lengthVerifier *length.Verifier
	{
		c := length.VerifierConfig{}

		lengthVerifier, err = length.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var stateVerifier *state.Verifier
	{
		c := state.VerifierConfig{}

		stateVerifier, err = state.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u := &Updater{
		logger: config.Logger,
		redigo: config.Redigo,
		rescue: config.Rescue,

		verify: []updater.Interface{
			emptyVerifier,
			lengthVerifier,
			stateVerifier,
		},
	}

	return u, nil
}
