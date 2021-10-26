package empty

import (
	"context"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"

	"github.com/venturemark/apicommon/pkg/metadata"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *texupd.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
		if req.Obj[0].Property == nil {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.VentureID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj[0].Property.Head == "" {
			return false, nil
		}
	}

	{
		// Creating text updates requires text or an attachment to be specified. It is not valid
		// to request the creation of text updates without providing any
		// information.
		if len(req.Obj[0].Property.Attachments) == 0 && req.Obj[0].Property.Text == "" {
			return false, nil
		}
	}

	{
		for _, attachment := range req.Obj[0].Property.Attachments {
			if attachment == nil {
				return false, nil
			}
			if attachment.Type == "" {
				return false, nil
			}
			if attachment.Addr == "" {
				return false, nil
			}
		}
	}

	return true, nil
}
