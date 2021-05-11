package time

import (
	"context"
	"strconv"
	"time"

	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
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

func (v *Verifier) Verify(ctx context.Context, req *message.UpdateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	var mei string
	{
		mei = req.Obj[0].Metadata[metadata.MessageID]
		if mei == "" {
			return false, nil
		}
	}

	{
		i, err := strconv.ParseInt(mei, 10, 64)
		if err != nil {
			return false, tracer.Mask(err)
		}

		now := v.now().UTC()
		uni := time.Unix(i, 0).Add(5 * time.Minute)

		// We allow messages to be modified for the first 5 minutes after their
		// creation. After this period we prohibit changes to messages. This is
		// to guarantee integrity of track records.
		if now.After(uni) {
			return false, nil
		}
	}

	return true, nil
}
