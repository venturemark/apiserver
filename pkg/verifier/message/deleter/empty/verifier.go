package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

// Verify checks if there is any information given for deleting messages. What
// we need is the organization ID, the message ID, the timeline ID and the
// update ID associated with the message.
func (v *Verifier) Verify(req *message.DeleteI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.OrganizationID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.MessageID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.UpdateID] == "" {
			return false, nil
		}
	}

	return true, nil
}
