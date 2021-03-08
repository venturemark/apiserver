package exist

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type VerifierConfig struct {
	Redigo redigo.Interface
}

type Verifier struct {
	redigo redigo.Interface
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	if config.Redigo == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Redigo must not be empty", config)
	}

	v := &Verifier{
		redigo: config.Redigo,
	}

	return v, nil
}

func (v *Verifier) Verify(req *message.DeleteI) (bool, error) {
	var err error

	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	var mei float64
	{
		mei, err = strconv.ParseFloat(req.Obj[0].Metadata[metadata.MessageID], 64)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	var tii string
	{
		tii = req.Obj[0].Metadata[metadata.TimelineID]
	}

	var upi string
	{
		upi = req.Obj[0].Metadata[metadata.UpdateID]
	}

	var vei string
	{
		vei = req.Obj[0].Metadata[metadata.VentureID]
	}

	{
		k := fmt.Sprintf(key.Message, vei, tii, upi)
		s := mei

		exi, err := v.redigo.Sorted().Exists().Score(k, s)
		if err != nil {
			return false, tracer.Mask(err)
		}

		if !exi {
			return false, nil
		}
	}

	return true, nil
}
