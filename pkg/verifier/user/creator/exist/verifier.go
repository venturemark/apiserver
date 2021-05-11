package exist

import (
	"context"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/association"
)

type VerifierConfig struct {
	Association *association.Association
}

type Verifier struct {
	association *association.Association
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	if config.Association == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Association must not be empty", config)
	}

	v := &Verifier{
		association: config.Association,
	}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *user.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	var clk *key.Key
	{
		clk = key.Claim(req.Obj[0].Metadata)
	}

	{
		exi, err := v.association.Exists(clk)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if exi {
			return false, nil
		}
	}

	return true, nil
}
