package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/texupd"

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

// Verify checks if there is any information given for creating text updates.
func (v *Verifier) Verify(req *texupd.CreateI) (bool, error) {
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
		if req.Obj.Metadata[metadata.AudienceID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
	}

	{
		// Creating text updates requires text to be specified. It is not valid
		// to request the creation of text updates without providing any
		// information.
		if req.Obj.Property.Text == "" {
			return false, nil
		}
	}

	return true, nil
}
