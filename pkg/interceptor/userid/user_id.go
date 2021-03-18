package userid

import (
	"context"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/association"
	"github.com/venturemark/apiserver/pkg/context/claimid"
	"github.com/venturemark/apiserver/pkg/context/userid"
)

type InterceptorConfig struct {
	Association *association.Association
	Logger      logger.Interface
}

type Interceptor struct {
	association *association.Association
	logger      logger.Interface
}

func NewInterceptor(config InterceptorConfig) (*Interceptor, error) {
	if config.Association == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Association must not be empty", config)
	}
	if config.Logger == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	e := &Interceptor{
		association: config.Association,
		logger:      config.Logger,
	}

	return e, nil
}

func (e *Interceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, inf *grpc.UnaryServerInfo, han grpc.UnaryHandler) (interface{}, error) {
		var clk *key.Key
		{
			cli, ok := claimid.FromContext(ctx)
			if !ok {
				return nil, tracer.Maskf(invalidMetadataError, "subject must not be empty")
			}

			met := map[string]string{
				metadata.ClaimID: cli,
			}

			clk = key.Claim(met)
		}

		{
			usi, err := e.association.Search(clk)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			ctx = userid.NewContext(ctx, usi)
		}

		return han(ctx, req)
	}
}
