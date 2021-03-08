package timeline

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

// Verify checks if the timeline which the text update should be created for
// does even exist. Creating text updates requires a timeline to exist.
func (v *Verifier) Verify(req *texupd.CreateI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	var tii string
	var vei string
	{
		tii = req.Obj[0].Metadata[metadata.TimelineID]
		vei = req.Obj[0].Metadata[metadata.VentureID]

		if tii == "" {
			return false, nil
		}
		if vei == "" {
			return false, nil
		}
	}

	{
		i, err := strconv.ParseFloat(tii, 64)
		if err != nil {
			return false, tracer.Mask(err)
		}

		k := fmt.Sprintf(key.Timeline, vei)
		s := i

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
