package daemon

import (
	"context"
	"net"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/client"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/handler"
	"github.com/venturemark/apiserver/pkg/handler/metric"
	"github.com/venturemark/apiserver/pkg/handler/metupd"
	"github.com/venturemark/apiserver/pkg/server/grpc"
	"github.com/venturemark/apiserver/pkg/storage"
)

type runner struct {
	flag   *flag
	logger logger.Interface
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

	var redisClient redigo.Interface
	{
		c := client.Config{
			Address: net.JoinHostPort(r.flag.Redis.Host, r.flag.Redis.Port),
		}

		redisClient, err = client.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var redisStorage *storage.Storage
	{
		c := storage.Config{
			Logger: r.logger,
			Redigo: redisClient,
		}

		redisStorage, err = storage.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var metricHandler *metric.Handler
	{
		c := metric.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		metricHandler, err = metric.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var metupdHandler *metupd.Handler
	{
		c := metupd.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		metupdHandler, err = metupd.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var g *grpc.Server
	{
		c := grpc.ServerConfig{
			Logger: r.logger,
			Handlers: []handler.Interface{
				metricHandler,
				metupdHandler,
			},

			Host: r.flag.ApiServer.Host,
			Port: r.flag.ApiServer.Port,
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
