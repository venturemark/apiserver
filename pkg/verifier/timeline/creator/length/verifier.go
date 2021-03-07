package length

import (
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *timeline.CreateI) (bool, error) {
	{
		if req.Obj == nil {
			return true, nil
		}
		if req.Obj.Property == nil {
			return true, nil
		}
	}

	{
		if len(req.Obj.Property.Desc) > 280 {
			return false, nil
		}
		if len(req.Obj.Property.Name) > 32 {
			return false, nil
		}
	}

	return true, nil
}
