package empty

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *invite.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return true, nil
		}
		if req.Obj[0].Metadata == nil {
			return true, nil
		}
		if req.Obj[0].Property == nil {
			return true, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.VentureID] == "" {
			return false, nil
		}
	}

	return true, nil
}
