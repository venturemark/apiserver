package empty

import (
	"github.com/venturemark/apigengo/pkg/pbf/timeline"

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

// Verify checks if there is any information given for deleting timelines. What
// we need is the audience ID and the timeline ID associated with the timeline.
func (v *Verifier) Verify(req *timeline.DeleteI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.AudienceID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
	}

	return true, nil
}
