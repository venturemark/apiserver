package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/metupd"

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

// Verify checks if there is any information given for creating metric updates.
// Succifient information for both of these resources must be provided.
func (v *Verifier) Verify(req *metupd.CreateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Property == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.UserID] == "" {
			return false, nil
		}
	}

	{
		// Creating metric updates requires the both of the resources to be
		// specified. It is not valid to request the creation of metric updates
		// without providing any information.
		if req.Obj.Property.Data == nil || req.Obj.Property.Text == "" {
			return false, nil
		}
	}

	{
		// If data is given but is in fact empty we consider this request to be
		// invalid.
		if req.Obj.Property.Data != nil && len(req.Obj.Property.Data) == 0 {
			return false, nil
		}
	}

	return true, nil
}
