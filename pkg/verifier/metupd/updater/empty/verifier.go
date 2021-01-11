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

// Verify checks if there is any information given for modifying metrics or
// updates. At least one of them must be updated with a request. Both of them
// can be updated at the same time.
func (v *Verifier) Verify(req *metupd.UpdateI) (bool, error) {
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
		met := req.Obj.Metadata[metadata.MetricID] == ""
		dat := req.Obj.Property.Data == nil

		upd := req.Obj.Metadata[metadata.UpdateID] == ""
		tex := req.Obj.Property.Text == ""

		// Ensure a metric ID if data is provided to be updated.
		if met && !dat || !met && dat {
			return false, nil
		}
		// Ensure an update ID if text is provided to be updated.
		if upd && !tex || !upd && tex {
			return false, nil
		}
		// Ensure that either data or text is provided to be updated.
		if dat && tex {
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
