package text

import (
	"context"
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/slate"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apicommon/pkg/metadata"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

func (v *Verifier) Verify(ctx context.Context, req *texupd.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Property == nil {
			return false, nil
		}
	}

	{
		updateFormat := req.Obj[0].Metadata[metadata.UpdateFormat]
		text := req.Obj[0].Property.Text
		var length int
		if updateFormat == "slate" {
			var nodes slate.Nodes
			err := json.Unmarshal([]byte(text), &nodes)
			if err != nil {
				return false, tracer.Mask(err)
			}

			length = slate.Length(nodes)
		} else {
			length = len(text)
		}

		if length > 280 {
			return false, nil
		}
	}

	return true, nil
}
