package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	DonCha   <-chan struct{}
	ErrCha   chan<- error
	GRPCHost string
	GRPCPort string
	HTTPHost string
	HTTPPort string
}

type Server struct {
	interceptor []grpc.UnaryServerInterceptor
	logger      logger.Interface
	handler     []handler.Interface

	donCha   <-chan struct{}
	errCha   chan<- error
	grpcHost string
	grpcPort string
	httpHost string
	httpPort string
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
	if config.GRPCHost == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.GRPCHost must not be empty", config)
	}
	if config.GRPCPort == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.GRPCPort must not be empty", config)
	}
	if config.HTTPHost == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.HTTPHost must not be empty", config)
	}
	if config.HTTPPort == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.HTTPPort must not be empty", config)
	}

	s := &Server{
		interceptor: config.Interceptor,
		logger:      config.Logger,
		handler:     config.Handler,

		donCha:   config.DonCha,
		errCha:   config.ErrCha,
		grpcHost: config.GRPCHost,
		grpcPort: config.GRPCPort,
		httpHost: config.HTTPHost,
		httpPort: config.HTTPPort,
	}

	return s, nil
}

func (s *Server) ListenGRPC() {
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
		l, err = net.Listen("tcp", net.JoinHostPort(s.grpcHost, s.grpcPort))
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}

	s.logger.Log(context.Background(), "level", "info", "message", fmt.Sprintf("grpc server running at %s", l.Addr().String()))

	{
		err = ser.Serve(l)
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}
}

func (s *Server) ListenHTTP() {
	a := net.JoinHostPort(s.httpHost, s.httpPort)

	{
		http.Handle("/metrics", promhttp.Handler())
	}

	s.logger.Log(context.Background(), "level", "info", "message", fmt.Sprintf("http server running at %s", a))

	{
		err := http.ListenAndServe(a, nil)
		if err != nil {
			s.errCha <- tracer.Mask(err)
		}
	}
}
