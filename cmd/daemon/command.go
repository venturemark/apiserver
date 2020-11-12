package daemon

import (
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/client"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/storage"
)

const (
	name  = "daemon"
	short = "Run the apiserver process and serve grpc traffic."
	long  = "Run the apiserver process and serve grpc traffic."
)

type Config struct {
	Logger logger.Interface
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var r redigo.Interface
	{
		c := client.Config{}

		r, err = client.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var s *storage.Storage
	{
		c := storage.Config{
			Logger: config.Logger,
			Redigo: r,
		}

		s, err = storage.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var c *cobra.Command
	{
		f := &flag{}

		r := &runner{
			flag:    f,
			logger:  config.Logger,
			storage: s,
		}

		c = &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
			RunE:  r.Run,
		}

		f.Init(c)
	}

	return c, nil
}
