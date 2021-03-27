package patch

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *message.UpdateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
		if len(req.Obj[0].Jsnpatch) == 0 {
			return false, nil
		}
	}

	{
		if req.Obj[0].Metadata[metadata.MessageID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.TimelineID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.UpdateID] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.VentureID] == "" {
			return false, nil
		}
	}

	{
		for _, j := range req.Obj[0].Jsnpatch {
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

	{
		for i := range req.Obj[0].Jsnpatch {
			if req.Obj[0].Jsnpatch[i].Pat != "/obj/property/text" {
				continue
			}

			valEmp := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == ""
			valLen := req.Obj[0].Jsnpatch[i].Val != nil && len(*req.Obj[0].Jsnpatch[i].Val) <= 280

			if !valEmp && !valLen {
				return false, nil
			}
		}
	}

	return true, nil
}
