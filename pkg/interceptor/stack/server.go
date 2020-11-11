package stack

import (
	"context"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
)

type InterceptorConfig struct {
	Logger logger.Interface
}

type Interceptor struct {
	logger logger.Interface
}

func NewInterceptor(config InterceptorConfig) (*Interceptor, error) {
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	e := &Interceptor{
		logger: config.Logger,
	}

	return e, nil
}

func (e *Interceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, inf *grpc.UnaryServerInfo, han grpc.UnaryHandler) (interface{}, error) {
		res, err := han(ctx, req)
		if err != nil {
			e.logger.Log(ctx, "level", "error", "message", "request failed", "stack", tracer.JSON(err))
		}

		return res, tracer.Mask(err)
	}
}
