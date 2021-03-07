package exist

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
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

func (v *Verifier) Verify(req *texupd.DeleteI) (bool, error) {
	var err error

	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	var upi float64
	{
		upi, err = strconv.ParseFloat(req.Obj.Metadata[metadata.UpdateID], 64)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	var tii string
	{
		tii = req.Obj.Metadata[metadata.TimelineID]
	}

	var vei string
	{
		vei = req.Obj.Metadata[metadata.VentureID]
	}

	{
		k := fmt.Sprintf(key.Update, vei, tii)
		s := upi

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
