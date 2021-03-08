package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
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
		if len(req.Obj) != 1 {
			return true, nil
		}
		if req.Obj[0].Metadata == nil {
			return true, nil
		}
		if req.Obj[0].Property == nil {
			return true, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.VentureID] == "" {
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
