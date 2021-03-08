package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

// Verify checks if there is any information given for creating text updates.
func (v *Verifier) Verify(req *texupd.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
		if req.Obj[0].Property == nil {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.UserID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.VentureID] == "" {
			return false, nil
		}
	}

	{
		// Creating text updates requires text to be specified. It is not valid
		// to request the creation of text updates without providing any
		// information.
		if req.Obj[0].Property.Text == "" {
			return false, nil
		}
	}

	return true, nil
}
