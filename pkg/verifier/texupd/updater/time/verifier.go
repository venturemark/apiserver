package time

import (
	"strconv"
	"time"

	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/metadata"
)

type VerifierConfig struct {
	Now func() time.Time
}

type Verifier struct {
	now func() time.Time
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	if config.Now == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Now must not be empty", config)
	}

	v := &Verifier{
		now: config.Now,
	}

	return v, nil
}

// Verify checks if a text update is too old to be modified. We have a
// theshold of 5 minutes after creation.
func (v *Verifier) Verify(req *texupd.UpdateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	var uid string
	{
		uid = req.Obj.Metadata[metadata.UpdateID]
		if uid == "" {
			return false, nil
		}
	}

	{
		i, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			return false, tracer.Mask(err)
		}

		now := v.now().UTC()
		uni := time.Unix(i, 0).Add(5 * time.Minute)

		// We allow text updates to be modified for the first 5 minutes after
		// their creation. After this period we prohibit changes to text
		// updates. This is to guarantee integrity of track records.
		if now.After(uni) {
			return false, nil
		}
	}

	return true, nil
}
