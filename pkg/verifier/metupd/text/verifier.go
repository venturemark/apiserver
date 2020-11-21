package text

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

// Verify checks for text to be valid, if provided with the update request. It
// is legitimate to not provide any text to be updated, if the update request
// contains information to modify the data of a metric update. This is then
// verified in other verifier implementations.
func (v *Verifier) Verify(req *metupd.UpdateI) (bool, error) {
	{
		// If no data is provided it may not be an update request to modify
		// text. It may only be an update request to modify data. This is then
		// handled by another verifier.
		if req.Obj == nil {
			return true, nil
		}
		if req.Obj.Property == nil {
			return true, nil
		}
	}

	{
		// Updating the text of metric updates is optional. One may just wish to
		// update datapoints of a metric update. If the update text is provided,
		// it is still limited to up to 280 characters. Nobody should be able to
		// update metric updates with longer text.
		if req.Obj.Property.Text != "" && len(req.Obj.Property.Text) > 280 {
			return false, nil
		}
	}

	return true, nil
}
