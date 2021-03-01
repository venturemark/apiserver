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
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	var oid string
	var tid string
	{
		oid = req.Obj.Metadata[metadata.OrganizationID]
		tid = req.Obj.Metadata[metadata.TimelineID]

		if tid == "" {
			return false, nil
		}
		if oid == "" {
			return false, nil
		}
	}

	{
		i, err := strconv.ParseFloat(tid, 64)
		if err != nil {
			return false, tracer.Mask(err)
		}

		k := fmt.Sprintf(key.TimelineResource, oid)
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
