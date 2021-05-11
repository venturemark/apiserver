package empty

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/user"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *user.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Property == nil {
			return false, nil
		}
	}

	{
		if req.Obj[0].Property.Name == "" {
			return false, nil
		}
	}

	return true, nil
}
