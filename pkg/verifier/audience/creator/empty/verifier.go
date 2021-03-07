package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *audience.CreateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
		if req.Obj.Property == nil {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.UserID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.VentureID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj.Property.Name == "" {
			return false, nil
		}
	}

	{
		if len(req.Obj.Property.Tmln) == 0 {
			return false, nil
		}
	}

	// {
	// 	if len(req.Obj.Property.User) == 0 {
	// 		return false, nil
	// 	}
	// }

	return true, nil
}
