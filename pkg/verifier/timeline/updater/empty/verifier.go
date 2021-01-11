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

// Verify checks if there is any information given for creating timelines. The
// only piece of information we need is the audience ID provided with the object
// metadata.
func (v *Verifier) Verify(req *timeline.UpdateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
		if req.Obj.Property == nil {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.AudienceID] == "" {
			return false, nil
		}
	}

	{
		if req.Obj.Property.Name == "" {
			return false, nil
		}
	}

	return true, nil
}
