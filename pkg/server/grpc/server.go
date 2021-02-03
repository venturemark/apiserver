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
	"github.com/venturemark/apiserver/pkg/interceptor/stack"
)

type ServerConfig struct {
	Logger   logger.Interface
	Handlers []handler.Interface

	Host string
	Port string
}

type Server struct {
	logger   logger.Interface
	handlers []handler.Interface

	server *grpc.Server

	host string
	port string
}

func NewServer(config ServerConfig) (*Server, error) {
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

	var e *stack.Interceptor
	{
		c := stack.InterceptorConfig{
			Logger: config.Logger,
		}

		e, err = stack.NewInterceptor(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var u []grpc.UnaryServerInterceptor
	{
		u = append(u, e.Interceptor())
	}

	s := &Server{
		logger:   config.Logger,
		handlers: config.Handlers,

		server: grpc.NewServer(
			grpc.ChainUnaryInterceptor(u...),
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

func (s *Server) Listen() error {
	var err error

	var l net.Listener
	{
		l, err = net.Listen("tcp", net.JoinHostPort(s.host, s.port))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	s.logger.Log(context.Background(), "level", "info", "message", fmt.Sprintf("server running at %s", l.Addr().String()))

	err = s.server.Serve(l)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}
