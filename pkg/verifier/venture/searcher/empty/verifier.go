package empty

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/venture"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *venture.SearchI) (bool, error) {
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
			ven := req.Obj[0].Metadata[metadata.VentureID]
			ves := req.Obj[0].Metadata["venture.venturemark.co/slug"]

			if sub == "" && ven == "" && ves == "" {
				return false, nil
			}
		}
	}

	return true, nil
}
