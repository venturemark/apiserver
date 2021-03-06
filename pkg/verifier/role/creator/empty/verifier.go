package empty

import (
	"context"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/role"
)

var (
	resources = []string{
		"invite",
		"message",
		"timeline",
		"update",
		"user",
		"venture",
	}

	roles = []string{
		"member",
		"owner",
		"reader",
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

func (v *Verifier) Verify(ctx context.Context, req *role.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	{
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
