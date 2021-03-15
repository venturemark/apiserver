package daemon

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/gateway"
	"github.com/venturemark/permission/pkg/ingress"
	"github.com/venturemark/permission/pkg/resolver"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/client"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/rescue/pkg/engine"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/handler"
	"github.com/venturemark/apiserver/pkg/handler/audience"
	"github.com/venturemark/apiserver/pkg/handler/message"
	"github.com/venturemark/apiserver/pkg/handler/role"
	"github.com/venturemark/apiserver/pkg/handler/texupd"
	"github.com/venturemark/apiserver/pkg/handler/timeline"
	"github.com/venturemark/apiserver/pkg/handler/update"
	"github.com/venturemark/apiserver/pkg/handler/user"
	"github.com/venturemark/apiserver/pkg/handler/venture"
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

	var redigoClient redigo.Interface
	{
		c := client.Config{
			Address: net.JoinHostPort(r.flag.Redis.Host, r.flag.Redis.Port),
			Kind:    r.flag.Redis.Kind,
		}

		redigoClient, err = client.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var rescueEngine rescue.Interface
	{
		c := engine.Config{
			Logger: r.logger,
			Redigo: redigoClient,
		}

		rescueEngine, err = engine.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var ingressGateway permission.Ingress
	{
		c := ingress.Config{}

		ingressGateway, err = ingress.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var resourceResolver permission.Resource
	{
		c := resolver.Config{
			Logger: r.logger,
			Redigo: redigoClient,
		}

		resourceResolver, err = resolver.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var permissionGateway permission.Gateway
	{
		c := gateway.Config{
			Ingress:  ingressGateway,
			Resource: resourceResolver,
		}

		permissionGateway, err = gateway.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var redisStorage *storage.Storage
	{
		c := storage.Config{
			Logger:     r.logger,
			Permission: permissionGateway,
			Redigo:     redigoClient,
			Rescue:     rescueEngine,
		}

		redisStorage, err = storage.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var audienceHandler *audience.Handler
	{
		c := audience.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		audienceHandler, err = audience.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var messageHandler *message.Handler
	{
		c := message.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		messageHandler, err = message.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var roleHandler *role.Handler
	{
		c := role.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		roleHandler, err = role.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var texupdHandler *texupd.Handler
	{
		c := texupd.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		texupdHandler, err = texupd.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var timelineHandler *timeline.Handler
	{
		c := timeline.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		timelineHandler, err = timeline.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var updateHandler *update.Handler
	{
		c := update.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		updateHandler, err = update.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var userHandler *user.Handler
	{
		c := user.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		userHandler, err = user.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var ventureHandler *venture.Handler
	{
		c := venture.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		ventureHandler, err = venture.NewHandler(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var donCha chan struct{}
	var errCha chan error
	var sigCha chan os.Signal
	{
		donCha = make(chan struct{})
		errCha = make(chan error, 1)
		sigCha = make(chan os.Signal, 2)

		defer close(donCha)
		defer close(errCha)
		defer close(sigCha)
	}

	var g *grpc.Server
	{
		c := grpc.ServerConfig{
			DonCha: donCha,
			ErrCha: errCha,
			Logger: r.logger,
			Handlers: []handler.Interface{
				audienceHandler,
				messageHandler,
				roleHandler,
				texupdHandler,
				timelineHandler,
				updateHandler,
				userHandler,
				ventureHandler,
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

		go g.Listen()
	}

	{
		signal.Notify(sigCha, os.Interrupt, syscall.SIGTERM)

		select {
		case err := <-errCha:
			return tracer.Mask(err)

		case <-sigCha:
			select {
			case <-time.After(r.flag.ApiServer.TerminationGracePeriod):
			case <-sigCha:
			}

			return nil
		}
	}
}
