package empty

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/audience"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *audience.UpdateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
		if len(req.Obj.Jsnpatch) == 0 {
			return false, nil
		}
	}

	{
		if req.Obj.Metadata[metadata.AudienceID] == "" {
			return false, nil
		}
		if req.Obj.Metadata[metadata.VentureID] == "" {
			return false, nil
		}
	}

	{
		for _, j := range req.Obj.Jsnpatch {
			opeAdd := j.Ope == "add"
			opeRem := j.Ope == "remove"
			opeRep := j.Ope == "replace"
			opeTes := j.Ope == "test"
			patEmp := j.Pat == ""
			valEmp := j.Val == nil

			if !opeAdd && !opeRem && !opeRep && !opeTes {
				return false, nil
			}
			if patEmp {
				return false, nil
			}
			if !opeRem && valEmp {
				return false, nil
			}
		}
	}

	return true, nil
}
