package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func (t *Timeline) Verify(req *metric.SearchI) (bool, error) {
	for _, v := range t.verify {
		ok, err := v.Verify(req)
		if err != nil {
			return false, tracer.Mask(err)
		}
		if !ok {
			return false, nil
		}
	}

	return true, nil
}
