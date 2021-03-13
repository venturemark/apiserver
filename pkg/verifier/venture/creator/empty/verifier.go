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

func (v *Verifier) Verify(req *venture.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata != nil {
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
