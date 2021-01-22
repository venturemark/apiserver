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

// Verify checks if there is any information given for creating timelines. We
// need the organization ID and timeline ID provided with the object metadata.
// We need at least one of the existing properties to be specified in order to
// update the specified timeline.
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
		if req.Obj.Metadata[metadata.OrganizationID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
	}

	{
		desEmp := req.Obj.Property.Desc == nil
		namEmp := req.Obj.Property.Name == nil
		staEmp := req.Obj.Property.Stat == nil

		if desEmp && namEmp && staEmp {
			return false, nil
		}
	}

	return true, nil
}
