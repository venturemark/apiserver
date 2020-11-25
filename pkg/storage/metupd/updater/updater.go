package updater

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater/consistency"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater/empty"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater/space"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater/text"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater/time"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/updater/value"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Updater struct {
	logger logger.Interface
	redigo redigo.Interface

	verify []updater.Interface
}

func New(config Config) (*Updater, error) {
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

	var emptyVerifier *empty.Verifier
	{
		c := empty.VerifierConfig{}

		emptyVerifier, err = empty.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var spaceVerifier *space.Verifier
	{
		c := space.VerifierConfig{}

		spaceVerifier, err = space.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var textVerifier *text.Verifier
	{
		c := text.VerifierConfig{}

		textVerifier, err = text.NewVerifier(c)
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

	var valueVerifier *value.Verifier
	{
		c := value.VerifierConfig{}

		valueVerifier, err = value.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u := &Updater{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []updater.Interface{
			consistencyVerifier,
			emptyVerifier,
			spaceVerifier,
			textVerifier,
			timeVerifier,
			valueVerifier,
		},
	}

	return u, nil
}
