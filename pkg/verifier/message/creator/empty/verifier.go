package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/message"

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

// Verify checks if there is any information given for creating messages. What
// we need is the text of the message and sufficient metadata for association.
func (v *Verifier) Verify(req *message.CreateI) (bool, error) {
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
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.UpdateID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.UserID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj.Property.Text == "" {
			return false, nil
		}
	}

	return true, nil
}
