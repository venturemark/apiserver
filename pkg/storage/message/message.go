package message

import (
	"github.com/venturemark/permission"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage/message/creator"
	"github.com/venturemark/apiserver/pkg/storage/message/deleter"
	"github.com/venturemark/apiserver/pkg/storage/message/searcher"
)

type Config struct {
	Permission permission.Gateway
	Logger     logger.Interface
	Redigo     redigo.Interface
	Rescue     rescue.Interface
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
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		cre, err = creator.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var del *deleter.Deleter
	{
		c := deleter.Config{
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
		}

		del, err = deleter.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sea *searcher.Searcher
	{
		c := searcher.Config{
			Permission: config.Permission,
			Logger:     config.Logger,
			Redigo:     config.Redigo,
			Rescue:     config.Rescue,
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
