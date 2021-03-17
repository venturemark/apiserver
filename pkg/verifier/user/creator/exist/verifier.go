package exist

import (
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/user"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/association"
)

type VerifierConfig struct {
	Association *association.Association
}

type Verifier struct {
	association *association.Association
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	if config.Association == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Association must not be empty", config)
	}

	v := &Verifier{
		association: config.Association,
	}

	return v, nil
}

func (v *Verifier) Verify(req *user.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	var suc string
	{
		suc = req.Obj[0].Metadata[metadata.SubjectClaim]
	}

	var suk *key.Key
	{
		met := map[string]string{
			metadata.ResourceKind: "subject",
			metadata.SubjectID:    suc,
		}

		suk = key.Subject(met)
	}

	{
		exi, err := v.association.Exists(suk)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if exi {
			return false, nil
		}
	}

	return true, nil
}
