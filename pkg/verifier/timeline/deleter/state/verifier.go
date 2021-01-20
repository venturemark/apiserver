package state

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/timeline/element"
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

// Verify checks the timeline requested to be deleted is in fact still in active
// state. It is only allowed to delete timelines which are archived.
func (v *Verifier) Verify(req *timeline.DeleteI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
	}

	var aid string
	{
		aid = req.Obj.Metadata[metadata.AudienceID]

		if aid == "" {
			return false, nil
		}
	}

	var tid float64
	{
		s := req.Obj.Metadata[metadata.TimelineID]
		if s == "" {
			return false, nil
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return false, tracer.Mask(err)
		}

		tid = f
	}

	{
		k := fmt.Sprintf(key.Timeline, aid)

		str, err := v.redigo.Sorted().Search().Score(k, tid, tid)
		if err != nil {
			return false, tracer.Mask(err)
		}

		_, _, _, s, err := element.Split(str[0])
		if err != nil {
			return false, tracer.Mask(err)
		}

		if s != "archived" {
			return false, nil
		}
	}

	return true, nil
}
