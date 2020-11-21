package space

import (
	"strings"

	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

const (
	// whitelist defines the possible allocations for dimensional spaces of data
	// bodies. Note that t is not included because the system provides this
	// dimensional space already for the unix timestamp of metric update
	// creation.
	whitelist = "abcdefghijklmnopqrsuvwxyz"
)

type VerifierConfig struct {
}

type Verifier struct {
}

func NewVerifier(config VerifierConfig) (*Verifier, error) {
	v := &Verifier{}

	return v, nil
}

// Verify checks for dimensional spaces to be valid, if any data is provided
// with the update request.
func (v *Verifier) Verify(req *metupd.UpdateI) (bool, error) {
	{
		// If no data is provided it may not be an update request to modify
		// data. It may only be an update request to modify text. This is then
		// handled by another verifier.
		if req.Obj == nil {
			return true, nil
		}
		if req.Obj.Property == nil {
			return true, nil
		}
		if req.Obj.Property.Data == nil {
			return true, nil
		}
	}

	{
		// Updating metrics is optional when updating metric updates. Somebody
		// may just wish to update their updates.
		if len(req.Obj.Property.Data) != 0 {
			// Dimensional spaces must be identified with single character
			// variables. Anything else other than e.g. x, y or z is invalid.
			// Additionally the reserved dimensional space t must also not be
			// supplied since the system provides that automatically according
			// to the unix timestamps of metric update creation.
			for _, d := range req.Obj.Property.Data {
				if len(d.Space) == 0 {
					return false, nil
				}
				if strings.Count(whitelist, d.Space) != 1 {
					return false, nil
				}
			}

			// We do not permit updating datapoints for the same dimensional
			// space twice. If the user tries to update a metric update with
			// e.g. the dimension y being duplicated, the request fails.
			m := map[string]int{}
			for _, d := range req.Obj.Property.Data {
				m[d.Space] = m[d.Space] + 1
			}
			for _, c := range m {
				if c != 1 {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
