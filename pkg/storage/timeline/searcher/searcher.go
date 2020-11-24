package searcher

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/verifier/timeline/search"
	"github.com/venturemark/apiserver/pkg/verifier/timeline/search/empty"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Searcher struct {
	logger logger.Interface
	redigo redigo.Interface

	verify []search.Interface
}

func New(config Config) (*Searcher, error) {
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

	s := &Searcher{
		logger: config.Logger,
		redigo: config.Redigo,

		verify: []search.Interface{
			emptyVerifier,
		},
	}

	return s, nil
}
