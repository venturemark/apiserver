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

func (v *Verifier) Verify(req *message.CreateI) (bool, error) {
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
		if req.Obj[0].Metadata[metadata.UpdateID] == "" {
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
		if req.Obj[0].Property.Text == "" {
			return false, nil
		}
	}

	return true, nil
}
