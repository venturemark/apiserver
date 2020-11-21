package value

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

// Verify checks for data to be consistent, if provided with the update request.
// It is legitimate to not provide any data to be updated, if the update request
// contains information to modify the text of a metric update. This is then
// verified in other verifier implementations.
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
		// If no data is provided it may not be an update request to modify
		// data. It may only be an update request to modify text. This is then
		// handled by another verifier.
		if req.Obj.Property == nil {
			return true, nil
		}
		if req.Obj.Property.Data == nil {
			return true, nil
		}
	}

	{
		// Updating metrics is optional when updating metric updates. Somebody
		// may just wish to update their updates.
		if len(req.Obj.Property.Data) != 0 {
			// We check for data values to be missing. If an update request is
			// meant to update data, data must be provided.
			for _, d := range req.Obj.Property.Data {
				if len(d.Value) == 0 {
					return false, nil
				}
			}

			// The amount of all datapoints must be equal across dimensions
			// provided. We do not permit inconsistencies within the request data.
			for i, d := range req.Obj.Property.Data {
				if i == 0 {
					continue
				}
				if len(req.Obj.Property.Data[0].Value) != len(d.Value) {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
