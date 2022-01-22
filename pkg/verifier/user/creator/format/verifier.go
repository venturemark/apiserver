package format

import (
	"context"

	"github.com/badoux/checkmail"
	"github.com/venturemark/apigengo/pkg/pbf/user"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func contains(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

var allowedPrepopulate = []string{"a", "b"}
var allowedSurvey = []string{"x", "y"}

func (v *Verifier) Verify(ctx context.Context, req *user.CreateI) (bool, error) {
	if prepopulate := req.Obj[0].Metadata["user.venturemark.co/prepopulate"]; !contains(prepopulate, allowedPrepopulate) {
		return false, nil
	}

	if survey := req.Obj[0].Metadata["user.venturemark.co/surveyResponse"]; !contains(survey, allowedSurvey) {
		return false, nil
	}

	{
		err := checkmail.ValidateFormat(req.Obj[0].Property.Mail)
		if err != nil {
			return false, nil
		}
	}

	return true, nil
}
