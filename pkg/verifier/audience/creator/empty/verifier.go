package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/audience"

	"github.com/venturemark/apiserver/pkg/metadata"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

// Verify checks if there is any information given for creating audiences. What
// we need is the name of the audience and a list of user IDs associated to it.
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
		if req.Obj.Metadata[metadata.OrganizationID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.UserID] == "" {
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
