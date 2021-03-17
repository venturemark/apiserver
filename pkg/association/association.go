package association

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/simple"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apicommon/pkg/key"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Association struct {
	logger logger.Interface
	redigo redigo.Interface
}

func New(config Config) (*Association, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	a := &Association{
		logger: config.Logger,
		redigo: config.Redigo,
	}

	return a, nil
}

func (a *Association) Create(suk *key.Key, usi string) error {
	var err error

	{
		k := suk.Elem()
		v := usi

		err = a.redigo.Simple().Create().Element(k, v)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (a *Association) Exists(suk *key.Key) (bool, error) {
	var err error

	var exi bool
	{
		k := suk.Elem()

		exi, err = a.redigo.Simple().Exists().Element(k)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return exi, nil
}

func (a *Association) Search(suk *key.Key) (string, error) {
	var err error

	var usi string
	{
		k := suk.Elem()

		usi, err = a.redigo.Simple().Search().Value(k)
		if simple.IsNotFound(err) {
			// fall through
		} else if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return usi, nil
}
