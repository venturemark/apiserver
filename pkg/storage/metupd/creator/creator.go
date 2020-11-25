package creator

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/metupd/creator"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/creator/consistency"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/creator/empty"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/creator/space"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/creator/text"
	"github.com/venturemark/apiserver/pkg/verifier/metupd/creator/value"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Creator struct {
	logger logger.Interface
	redigo redigo.Interface

	verify []creator.Interface
}

func New(config Config) (*Creator, error) {
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

	var emptyVerifier *empty.Verifier
	{
		c := empty.VerifierConfig{}

		emptyVerifier, err = empty.NewVerifier(c)
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

	c := &Creator{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []creator.Interface{
			consistencyVerifier,
			emptyVerifier,
			spaceVerifier,
			textVerifier,
			valueVerifier,
		},
	}

	return c, nil
}
