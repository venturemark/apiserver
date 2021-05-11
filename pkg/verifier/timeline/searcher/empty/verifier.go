package empty

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
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

func (v *Verifier) Verify(ctx context.Context, req *timeline.SearchI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	{
		suiEmp := req.Obj[0].Metadata[metadata.SubjectID] == ""
		tiiEmp := req.Obj[0].Metadata[metadata.TimelineID] == ""
		veiEmp := req.Obj[0].Metadata[metadata.VentureID] == ""

		if suiEmp && tiiEmp && veiEmp {
			return false, nil
		}

		if !suiEmp && (!tiiEmp || !veiEmp) {
			return false, nil
		}

		if !tiiEmp && veiEmp {
			return false, nil
		}

		if !tiiEmp && !suiEmp {
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
