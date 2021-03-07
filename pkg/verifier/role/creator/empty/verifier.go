package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/role"
)

var (
	resources = []string{
		"audience",
		"message",
		"timeline",
		"update",
		"venture",
	}

	roles = []string{
		"member",
		"owner",
	}
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *role.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.ResourceID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.ResourceKind] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.RoleKind] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.SubjectID] == "" {
			return false, nil
		}
	}

	{
		l := resources
		s := req.Obj[0].Metadata[metadata.ResourceKind]

		if !contains(l, s) {
			return false, nil
		}
	}

	{
		l := roles
		s := req.Obj[0].Metadata[metadata.RoleKind]

		if !contains(l, s) {
			return false, nil
		}
	}

	return true, nil
}

func contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}

	return false
}
