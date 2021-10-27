package text

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apicommon/pkg/slate"
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
		if len(req.Obj[0].Property.Head) > 280 {
			return false, nil
		}
	}

	if len(req.Obj[0].Property.Text) > 0 {
		var nodes slate.Nodes
		err := json.Unmarshal([]byte(req.Obj[0].Property.Text), &nodes)
		if err != nil {
			return false, tracer.Mask(err)
		}

		if slate.Length(nodes) > 600 {
			return false, nil
		}
	}

	{
		if len(req.Obj[0].Property.Attachments) > 1 {
			return false, nil
		}

		for _, attachment := range req.Obj[0].Property.Attachments {
			if attachment.Type != "image" {
				return false, nil
			}
			if len(attachment.Addr) > 200 {
				return false, nil
			}
			address, err := url.Parse(attachment.Addr)
			if err != nil {
				return false, nil
			}
			if address.Host != "res.cloudinary.com" {
				return false, nil
			}
		}
	}

	return true, nil
}
