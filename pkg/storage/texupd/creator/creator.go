package creator

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/texupd/creator"
	"github.com/venturemark/apiserver/pkg/verifier/texupd/creator/empty"
	"github.com/venturemark/apiserver/pkg/verifier/texupd/creator/text"
	"github.com/venturemark/apiserver/pkg/verifier/texupd/creator/timeline"
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

	var emptyVerifier *empty.Verifier
	{
		c := empty.VerifierConfig{}

		emptyVerifier, err = empty.NewVerifier(c)
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

	var timelineVerifier *timeline.Verifier
	{
		c := timeline.VerifierConfig{
			Redigo: config.Redigo,
		}

		timelineVerifier, err = timeline.NewVerifier(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	c := &Creator{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []creator.Interface{
			emptyVerifier,
			textVerifier,
			timelineVerifier,
		},
	}

	return c, nil
}