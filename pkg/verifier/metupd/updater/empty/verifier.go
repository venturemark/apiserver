package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

// Verify checks if there is any information given for modifying metrics or
// updates. At least one of them must be updated with a request. Both of them
// can be updated at the same time.
func (v *Verifier) Verify(req *metupd.UpdateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Property == nil {
			return false, nil
		}
	}

	{
		// Updating metric updates requires at least one of the resources to be
		// specified for modification. It is not valid to request the update of
		// any resource without providing any of these resources.
		if req.Obj.Property.Data == nil && req.Obj.Property.Text == "" {
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
