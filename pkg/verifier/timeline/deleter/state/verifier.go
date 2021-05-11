package state

import (
	"context"
	"encoding/json"

	"github.com/venturemark/apicommon/pkg/key"
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

func (v *Verifier) Verify(ctx context.Context, req *timeline.DeleteI) (bool, error) {
	{
		if len(req.Obj) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata == nil {
			return false, nil
		}
	}

	var tik *key.Key
	{
		tik = key.Timeline(req.Obj[0].Metadata)
	}

	{
		k := tik.List()
		s, err := v.redigo.Sorted().Search().Score(k, tik.ID().F(), tik.ID().F())
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
