package server

import (
	"context"
	"fmt"
	"net"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/venturemark/apiserver/pkg/handler"
)

type Config struct {
	Interceptor []grpc.UnaryServerInterceptor
	Logger      logger.Interface
	Handler     []handler.Interface

	DonCha <-chan struct{}
	ErrCha chan<- error
	Host   string
	Port   string
}

type Server struct {
	interceptor []grpc.UnaryServerInterceptor
	logger      logger.Interface
	handler     []handler.Interface

	donCha <-chan struct{}
	errCha chan<- error
	host   string
	port   string
}

func New(config Config) (*Server, error) {
	if len(config.Interceptor) == 0 {
		return nil, tracer.Maskf(invalidConfigError, "%T.Interceptor must not be empty", config)
	}
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if len(config.Handler) == 0 {
		return nil, tracer.Maskf(invalidConfigError, "%T.Handler must not be empty", config)
	}

	if config.DonCha == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.DonCha must not be empty", config)
	}
	if config.ErrCha == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.ErrCha must not be empty", config)
	}
	if config.Host == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Host must not be empty", config)
	}
	if config.Port == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Port must not be empty", config)
	}

	s := &Server{
		interceptor: config.Interceptor,
		logger:      config.Logger,
		handler:     config.Handler,

		donCha: config.DonCha,
		errCha: config.ErrCha,
		host:   config.Host,
		port:   config.Port,
	}

	return s, nil
}

func (s *Server) Listen() {
	var err error

	var ser *grpc.Server
	{
		ser = grpc.NewServer(grpc.ChainUnaryInterceptor(s.interceptor...))

		reflection.Register(ser)

		for _, h := range s.handler {
			h.Attach(ser)
		}
	}

	var l net.Listener
	{
		l, err = net.Listen("tcp", net.JoinHostPort(s.host, s.port))
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}

	s.logger.Log(context.Background(), "level", "info", "message", fmt.Sprintf("server running at %s", l.Addr().String()))

	{
		err = ser.Serve(l)
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}
}
