package format

import (
	"context"

	"github.com/badoux/checkmail"
	"github.com/venturemark/apigengo/pkg/pbf/user"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *user.CreateI) (bool, error) {
	{
		err := checkmail.ValidateFormat(req.Obj[0].Property.Mail)
		if err != nil {
			return false, nil
		}
	}

	return true, nil
}
