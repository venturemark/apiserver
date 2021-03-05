package state

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
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

	var vid string
	{
		vid = req.Obj.Metadata[metadata.VentureID]

		if vid == "" {
			return false, nil
		}
	}

	{
		k := fmt.Sprintf(key.Timeline, vid)

		s, err := v.redigo.Sorted().Search().Score(k, tid, tid)
		if err != nil {
			return false, tracer.Mask(err)
		}

		tim := &schema.Timeline{}
		err = json.Unmarshal([]byte(s[0]), tim)
		if err != nil {
			return false, tracer.Mask(err)
		}

		if tim.Obj.Property.Stat != "archived" {
			return false, nil
		}
	}

	return true, nil
}
