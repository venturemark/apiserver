package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/venturemark/apiserver/pkg/handler"
	"github.com/venturemark/apiserver/pkg/interceptor/oauth"
	"github.com/venturemark/apiserver/pkg/interceptor/stack"
)

type ServerConfig struct {
	DonCha   <-chan struct{}
	ErrCha   chan<- error
	Logger   logger.Interface
	Handlers []handler.Interface

	Host string
	Port string
}

type Server struct {
	donCha   <-chan struct{}
	errCha   chan<- error
	logger   logger.Interface
	handlers []handler.Interface

	server *grpc.Server

	host string
	port string
}

func NewServer(config ServerConfig) (*Server, error) {
	if config.DonCha == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.DonCha must not be empty", config)
	}
	if config.ErrCha == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.ErrCha must not be empty", config)
	}
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if len(config.Handlers) == 0 {
		return nil, tracer.Maskf(invalidConfigError, "%T.Handlers must not be empty", config)
	}

	if config.Host == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Host must not be empty", config)
	}
	if config.Port == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Port must not be empty", config)
	}

	var err error

	var oau *oauth.Interceptor
	{
		c := oauth.InterceptorConfig{
			Logger: config.Logger,
		}

		oau, err = oauth.NewInterceptor(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var sta *stack.Interceptor
	{
		c := stack.InterceptorConfig{
			Logger: config.Logger,
		}

		sta, err = stack.NewInterceptor(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cui []grpc.UnaryServerInterceptor
	{
		cui = append(cui, sta.Interceptor()) // stack first for error logging
		cui = append(cui, oau.Interceptor())
	}

	s := &Server{
		donCha:   config.DonCha,
		errCha:   config.ErrCha,
		logger:   config.Logger,
		handlers: config.Handlers,

		server: grpc.NewServer(
			grpc.ChainUnaryInterceptor(cui...),
		),

		host: config.Host,
		port: config.Port,
	}

	return s, nil
}

func (s *Server) Attach() error {
	reflection.Register(s.server)

	for _, h := range s.handlers {
		h.Attach(s.server)
	}

	return nil
}

func (s *Server) Listen() {
	var err error

	var l net.Listener
	{
		l, err = net.Listen("tcp", net.JoinHostPort(s.host, s.port))
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}

	s.logger.Log(context.Background(), "level", "info", "message", fmt.Sprintf("server running at %s", l.Addr().String()))

	{
		err = s.server.Serve(l)
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}
}
