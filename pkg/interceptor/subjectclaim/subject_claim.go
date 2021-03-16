package subjectclaim

import (
	"context"
	"crypto/sha256"
	"fmt"
	"regexp"

	"github.com/lestrrat-go/jwx/jwt"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/venturemark/apiserver/pkg/context/subjectclaim"
)

var (
	bearerScheme = regexp.MustCompile(`(?i)bearer `)
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
		var ok bool

		var m metadata.MD
		{
			m, ok = metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, tracer.Maskf(invalidMetadataError, "metadata must not be empty")
			}
		}

		var a string
		{
			l, ok := m["authorization"]
			if !ok {
				return nil, tracer.Maskf(invalidMetadataError, "authorization must not be empty")
			}

			if len(l) != 1 {
				return nil, tracer.Maskf(invalidMetadataError, "authorization must be given once")
			}

			a = l[0]

			if !bearerScheme.MatchString(a) {
				return nil, tracer.Maskf(invalidMetadataError, "authorization must be bearer scheme")
			}

			a = a[7:]
		}

		var sui string
		{
			t, err := jwt.ParseString(a)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			h := sha256.New()
			_, err = h.Write([]byte(t.Subject()))
			if err != nil {
				return nil, tracer.Mask(err)
			}
			sui = fmt.Sprintf("%x", h.Sum(nil))
		}

		{
			ctx = subjectclaim.NewContext(ctx, sui)
		}

		return han(ctx, req)
	}
}
