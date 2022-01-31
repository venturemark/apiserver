package format

import (
	"context"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/content"

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

// Keep in sync with https://github.com/venturemark/webclient/blob/767d411768e3e4fd45b42f16c8ded208233d7698/src/component/OnboardingGroup.tsx#L100-L103
var allowedSurvey = []string{
	"choicePeople",
	"choiceSearch",
	"choiceSocial",
	"choiceOther",
}

func (v *Verifier) Verify(ctx context.Context, req *user.CreateI) (bool, error) {
	var allowedPrepopulate []string
	for key := range content.DefaultTimelinesMap {
		allowedPrepopulate = append(allowedPrepopulate, key)
	}

	prepopulate := req.Obj[0].Metadata[metadata.UserPrepopulate]
	if prepopulate != "" && !contains(prepopulate, allowedPrepopulate) {
		return false, nil
	}

	survey := req.Obj[0].Metadata[metadata.UserSurveyResponse]
	if survey != "" && !contains(survey, allowedSurvey) {
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
