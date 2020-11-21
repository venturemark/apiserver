package consistency

import (
	"fmt"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apiserver/pkg/key"
	"github.com/venturemark/apiserver/pkg/metadata"
	"github.com/venturemark/apiserver/pkg/value/metric/timeline/data"
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

// Verify checks if a metric update is consistent with the data already recorded
// for a timeline. Tracking 3 datapoints across all dimensional spaces means
// that no 2 or no 4 datapoints can be provided with the update request, since
// this would lead to data inconsistencies.
func (v *Verifier) Verify(req *metupd.UpdateI) (bool, error) {
	{
		if req.Obj == nil {
			return false, nil
		}
		if req.Obj.Metadata == nil {
			return false, nil
		}
		if req.Obj.Property == nil {
			return false, nil
		}
		if req.Obj.Property.Data == nil {
			return false, nil
		}
	}

	{
		// Updating metric updates requires a timeline ID to be provided with
		// which the metric and the update can be associated with. If the
		// timeline ID is empty, we decline service for this request.
		if req.Obj.Metadata[metadata.Timeline] == "" {
			return false, nil
		}
	}

	{
		// Updating metrics is optional when updating metric updates. Somebody
		// may just wish to update their updates.
		if len(req.Obj.Property.Data) != 0 {
			// We always check the latest item of the sorted set to check the
			// amount of datapoints on the y axis. Due to this very check the
			// consistency of the sorted set is ensured, which means that
			// looking up a single element of the sorted set is sufficient.
			k := fmt.Sprintf(key.TimelineMetric, req.Obj.Metadata[metadata.Timeline])
			s, err := v.redigo.Scored().Search(k, 0, 1)
			if err != nil {
				return false, tracer.Mask(err)
			}
			if len(s) == 1 {
				_, val, err := data.Split(s[0])
				if err != nil {
					return false, tracer.Mask(err)
				}

				c := len(val[0].GetValue())
				y := len(req.Obj.Property.Data[0].Value)
				if c != y {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
