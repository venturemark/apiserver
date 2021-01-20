package state

import (
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

// Verify checks if the timeline state, if given, is either active or archived.
func (v *Verifier) Verify(req *timeline.UpdateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Property == nil {
			return false, nil
		}
	}

	{
		staEmp := req.Obj.Property.Stat != nil && *req.Obj.Property.Stat == ""
		staAct := req.Obj.Property.Stat != nil && *req.Obj.Property.Stat == "active"
		staArc := req.Obj.Property.Stat != nil && *req.Obj.Property.Stat == "archived"

		if !staEmp && !staAct && !staArc {
			return false, nil
		}
	}

	return true, nil
}
