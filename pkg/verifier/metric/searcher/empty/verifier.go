package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"

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

// Verify checks if there is any information given for searching metrics. The
// information we need is the audience ID and the timeline ID provided with the
// object metadata. It is only allowed to provide one search object since no
// more complex search queries are implemented yet.
func (v *Verifier) Verify(req *metric.SearchI) (bool, error) {
	{
		if req.Api != nil {
			return false, nil
		}
	}

	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.AudienceID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.TimelineID] == "" {
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
