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

// Verify checks if there is any information given for deleting audiences. What
// we need is the audience ID and the user ID associated with the audience.
func (v *Verifier) Verify(req *audience.DeleteI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.AudienceID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.UserID] == "" {
			return false, nil
		}
	}

	return true, nil
}
