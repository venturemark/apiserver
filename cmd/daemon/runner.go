package daemon

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/spf13/cobra"
	"github.com/venturemark/permission"
	"github.com/venturemark/permission/pkg/gateway"
	"github.com/venturemark/permission/pkg/ingress"
	"github.com/venturemark/permission/pkg/resolver"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/client"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/rescue/pkg/collector"
	"github.com/xh3b4sd/rescue/pkg/engine"
	"github.com/xh3b4sd/rescue/pkg/metric"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"

	"github.com/venturemark/apiserver/pkg/association"
	"github.com/venturemark/apiserver/pkg/handler"
	"github.com/venturemark/apiserver/pkg/handler/invite"
	"github.com/venturemark/apiserver/pkg/handler/message"
	"github.com/venturemark/apiserver/pkg/handler/role"
	"github.com/venturemark/apiserver/pkg/handler/texupd"
	"github.com/venturemark/apiserver/pkg/handler/timeline"
	"github.com/venturemark/apiserver/pkg/handler/update"
	"github.com/venturemark/apiserver/pkg/handler/user"
	"github.com/venturemark/apiserver/pkg/handler/venture"
	"github.com/venturemark/apiserver/pkg/interceptor/claimid"
	"github.com/venturemark/apiserver/pkg/interceptor/stack"
	"github.com/venturemark/apiserver/pkg/interceptor/userid"
	"github.com/venturemark/apiserver/pkg/server"
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

	//************************************************************************//

	var rescueMetric *metric.Collection
	{
		rescueMetric = metric.New()
	}

	var rescueCollector *collector.Collector
	{
		c := collector.Config{
			Logger: r.logger,
			Metric: rescueMetric,
		}

		rescueCollector, err = collector.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var rescueEngine rescue.Interface
	{
		c := engine.Config{
			Logger: r.logger,
			Metric: rescueMetric,
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

	var resourceResolver permission.Resolver
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
			Resolver: resourceResolver,
		}

		permissionGateway, err = gateway.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var associationMapper *association.Association
	{
		c := association.Config{
			Logger: r.logger,
			Redigo: redigoClient,
		}

		associationMapper, err = association.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var redisStorage *storage.Storage
	{
		c := storage.Config{
			Association: associationMapper,
			Logger:      r.logger,
			Permission:  permissionGateway,
			Redigo:      redigoClient,
			Rescue:      rescueEngine,
		}

		redisStorage, err = storage.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	//************************************************************************//

	var inviteHandler *invite.Handler
	{
		c := invite.HandlerConfig{
			Logger:  r.logger,
			Storage: redisStorage,
		}

		inviteHandler, err = invite.NewHandler(c)
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

	var sta *stack.Interceptor
	{
		c := stack.InterceptorConfig{
			Logger: r.logger,
		}

		sta, err = stack.NewInterceptor(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var sub *claimid.Interceptor
	{
		c := claimid.InterceptorConfig{
			Logger: r.logger,
		}

		sub, err = claimid.NewInterceptor(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var use *userid.Interceptor
	{
		c := userid.InterceptorConfig{
			Association: associationMapper,
			Logger:      r.logger,
		}

		use, err = userid.NewInterceptor(c)
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

	var g *server.Server
	{
		c := server.Config{
			Collector: []prometheus.Collector{
				collectors.NewGoCollector(),
				collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
				rescueCollector,
			},
			Interceptor: []grpc.UnaryServerInterceptor{
				sta.Interceptor(), // stack interceptor first for error logging
				sub.Interceptor(),
				use.Interceptor(), // user ID interceptor after subject claim interceptor for identity mapping
			},
			Logger: r.logger,
			Handler: []handler.Interface{
				inviteHandler,
				messageHandler,
				roleHandler,
				texupdHandler,
				timelineHandler,
				updateHandler,
				userHandler,
				ventureHandler,
			},

			DonCha:   donCha,
			ErrCha:   errCha,
			GRPCHost: r.flag.ApiServer.Host,
			GRPCPort: r.flag.ApiServer.Port,
			HTTPHost: r.flag.Metrics.Host,
			HTTPPort: r.flag.Metrics.Port,
		}

		g, err = server.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		go g.ListenGRPC()
		go g.ListenHTTP()
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
