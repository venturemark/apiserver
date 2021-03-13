package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/venture"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *venture.SearchI) (bool, error) {
	{
		if req.Api != nil {
			return false, nil
		}
	}

	{
		if len(req.Obj) != 0 {
			return false, nil
		}
	}

	return true, nil
}
