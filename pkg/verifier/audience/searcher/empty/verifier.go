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

// Verify checks if there is any information given for searching audiences.
func (v *Verifier) Verify(req *audience.SearchI) (bool, error) {
	{
		// We need a single object with a single metadata label for the user's
		// audience in order to fullfil the search request. We will extend
		// functionality here later.
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.OrganizationID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.UserID] == "" {
			return false, nil
		}
	}

	{
		// Any search request with object property specifics is not valid for
		// search requests at this point. We will extend functionality here
		// later.
		if req.Obj[0].Property != nil {
			return false, nil
		}
	}

	return true, nil
}