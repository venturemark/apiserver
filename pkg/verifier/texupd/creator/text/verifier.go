package text

import (
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

func (v *Verifier) Verify(req *texupd.CreateI) (bool, error) {
	{
		if req.Obj == nil {
			return true, nil
		}
		if req.Obj.Property == nil {
			return true, nil
		}
	}

	{
		// Updating the text of text updates must not exceed the character limit
		// of 280. Nobody should be able to create text updates with longer
		// text.
		if len(req.Obj.Property.Text) > 280 {
			return false, nil
		}
	}

	return true, nil
}
