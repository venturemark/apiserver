package patch

import (
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/invite"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(req *invite.UpdateI) (bool, error) {
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
		if req.Obj[0].Metadata[metadata.InviteID] == "" {
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
			if req.Obj[0].Jsnpatch[i].Pat != "/obj/property/mail" {
				continue
			}

			valEmp := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == ""
			valLen := req.Obj[0].Jsnpatch[i].Val != nil && len(*req.Obj[0].Jsnpatch[i].Val) <= 280

			if !valEmp && !valLen {
				return false, nil
			}
		}
	}

	{
		for i := range req.Obj[0].Jsnpatch {
			if req.Obj[0].Jsnpatch[i].Pat != "/obj/property/stat" {
				continue
			}

			valEmp := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == ""
			valPen := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "pending"
			valAcc := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "accepted"
			valRej := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "rejected"

			if !valEmp && !valPen && !valAcc && !valRej {
				return false, nil
			}
		}
	}

	return true, nil
}
