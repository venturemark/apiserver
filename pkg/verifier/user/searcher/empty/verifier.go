package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
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

func (v *Verifier) Verify(req *user.SearchI) (bool, error) {
	{
		if req.Api != nil {
			return false, nil
		}
	}

	{
		if len(req.Obj) > 1 {
			return false, nil
		}
	}

	{
		if len(req.Obj) == 1 {
			sub := req.Obj[0].Metadata[metadata.SubjectID]
			use := req.Obj[0].Metadata[metadata.UserID]

			if sub == "" && use == "" {
				return false, nil
			}
		}
	}

	return true, nil
}
