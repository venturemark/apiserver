package patch

import (
	"context"
	"github.com/venturemark/permission/pkg/label/visibility"
	"strings"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

type VerifierConfig struct {
}

type Verifier struct {
}

const metadataPath = "/obj/property/metatdata/"

var allowedVisibilities = []string{
	visibility.Private.Label(),
	visibility.Member.Label(),
	visibility.Public.Label(),
}

var allowedPermissionModels = []string{
	"writer",
	"reader",
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// https://datatracker.ietf.org/doc/html/rfc6901#section-3
func escapePatchPath(s string) string {
	return strings.ReplaceAll(s, "/", "~1")
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *timeline.UpdateI) (bool, error) {
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
		if req.Obj[0].Metadata[metadata.TimelineID] == "" {
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
			if req.Obj[0].Jsnpatch[i].Pat != "/obj/property/desc" {
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
			if req.Obj[0].Jsnpatch[i].Pat != "/obj/property/name" {
				continue
			}

			valEmp := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == ""
			valLen := req.Obj[0].Jsnpatch[i].Val != nil && len(*req.Obj[0].Jsnpatch[i].Val) <= 32

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
			valAct := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "active"
			valArc := req.Obj[0].Jsnpatch[i].Val != nil && *req.Obj[0].Jsnpatch[i].Val == "archived"

			if !valEmp && !valAct && !valArc {
				return false, nil
			}
		}
	}

	{
		for i := range req.Obj[0].Jsnpatch {
			if !strings.HasPrefix(req.Obj[0].Jsnpatch[i].Pat, metadataPath) {
				continue
			}

			var allowedValues []string
			switch strings.TrimPrefix(req.Obj[0].Jsnpatch[i].Pat, metadataPath) {
			case escapePatchPath(metadata.PermissionModel):
				allowedValues = allowedPermissionModels
			case escapePatchPath(metadata.ResourceVisibility):
				allowedValues = allowedVisibilities
			default:
				// Don't allow other metadata fields to be patched
				return false, nil
			}

			if contains([]string{"add", "replace"}, req.Obj[0].Jsnpatch[i].Ope) {
				if val := req.Obj[0].Jsnpatch[i].Val; val == nil {
					// Value must be defined for these operations
					return false, nil
				} else if !contains(allowedValues, *val) {
					// Value must be one of the allowed values
					return false, nil
				}
			}
		}
	}

	return true, nil
}
