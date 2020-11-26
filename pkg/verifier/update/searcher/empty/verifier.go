package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/update"

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

// Verify checks if there is any information given for searching updates. The
// information we need is the user ID and the timeline ID provided with the
// object metadata. It is only allowed to provide one search object since no
// more complex search queries are implemented yet.
func (v *Verifier) Verify(req *update.SearchI) (bool, error) {
	{
		if req.Api != nil {
			return false, nil
		}
	}

	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if len(req.Obj[0].Metadata) != 2 {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.UserID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj[0].Property != nil {
			return false, nil
		}
	}

	return true, nil
}
