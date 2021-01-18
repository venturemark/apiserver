package message

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/message/creator"
	"github.com/venturemark/apiserver/pkg/storage/message/deleter"
	"github.com/venturemark/apiserver/pkg/storage/message/searcher"
)

type Config struct {
	Logger logger.Interface
	Redigo redigo.Interface
}

type Message struct {
	Creator  *creator.Creator
	Deleter  *deleter.Deleter
	Searcher *searcher.Searcher
}

func New(config Config) (*Message, error) {
	var err error

	var cre *creator.Creator
	{
		c := creator.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		cre, err = creator.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var del *deleter.Deleter
	{
		c := deleter.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		del, err = deleter.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sea *searcher.Searcher
	{
		c := searcher.Config{
			Logger: config.Logger,
			Redigo: config.Redigo,
		}

		sea, err = searcher.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	m := &Message{
		Creator:  cre,
		Deleter:  del,
		Searcher: sea,
	}

	return m, nil
}
