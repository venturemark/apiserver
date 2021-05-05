package patch

import (
	"github.com/badoux/checkmail"
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
		if req.Obj[0].Metadata[metadata.ResourceKind] == "" {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.RoleKind] == "" {
			return false, nil
		}
	}

	{
		timEmp := req.Obj[0].Metadata[metadata.TimelineID] == ""
		venEmp := req.Obj[0].Metadata[metadata.VentureID] == ""
		timKin := req.Obj[0].Metadata[metadata.ResourceKind] == "timeline"
		venKin := req.Obj[0].Metadata[metadata.ResourceKind] == "venture"

		if timKin && timEmp && venEmp {
			return false, nil
		}
		if venKin && venEmp {
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
			if !valEmp {
				return false, nil
			}

			err := checkmail.ValidateFormat(*req.Obj[0].Jsnpatch[i].Val)
			if err != nil {
				return false, nil
			}
		}
	}

	{
		for i := range req.Obj[0].Jsnpatch {
			if req.Obj[0].Jsnpatch[i].Pat != "/obj/property/stat" {
				continue
			}

			incEmp := req.Obj[0].Metadata[metadata.InviteCode] == ""
			subEmp := req.Obj[0].Metadata[metadata.SubjectID] == ""
			valEmp := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == ""
			valPen := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "pending"
			valAcc := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "accepted"
			valRej := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "rejected"

			// We need the invite code in case the invite shall be accepted.
			// This is to verify the request's authenticity.
			if incEmp && valAcc {
				return false, nil
			}

			// We need the subject ID in case the invite shall be accepted. This
			// is to create a role making the associated user a member of the
			// venture for which the invite got accepted.
			if subEmp && valAcc {
				return false, nil
			}

			if !valEmp && !valPen && !valAcc && !valRej {
				return false, nil
			}
		}
	}

	return true, nil
}
