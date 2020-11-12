package daemon

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/handler"
	"github.com/venturemark/apiserver/pkg/handler/metric"
	"github.com/venturemark/apiserver/pkg/server/grpc"
	"github.com/venturemark/apiserver/pkg/storage"
)

type runner struct {
	flag    *flag
	logger  logger.Interface
	storage *storage.Storage
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return tracer.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var m *metric.Handler
	{
		c := metric.HandlerConfig{
			Logger:  r.logger,
			Storage: r.storage,
		}

		m, err = metric.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var g *grpc.Server
	{
		c := grpc.ServerConfig{
			Logger: r.logger,
			Handlers: []handler.Interface{
				m,
			},

			Host: r.flag.Host,
			Port: r.flag.Port,
		}

		g, err = grpc.NewServer(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = g.Attach()
		if err != nil {
			return tracer.Mask(err)
		}

		err = g.Listen()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
