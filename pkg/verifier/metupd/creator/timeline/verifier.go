package timeline

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
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

// Verify checks if the timeline which the metric update should be created for
// does even exist. Creating metric updates requires a timeline to exist.
func (v *Verifier) Verify(req *metupd.CreateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	var aid string
	var tid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]
		tid = req.Obj.Metadata[metadata.TimelineID]

		if aid == "" {
			return false, nil
		}
		if tid == "" {
			return false, nil
		}
	}

	{
		i, err := strconv.ParseFloat(tid, 64)
		if err != nil {
			return false, tracer.Mask(err)
		}

		k := fmt.Sprintf(key.Timeline, aid)
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
