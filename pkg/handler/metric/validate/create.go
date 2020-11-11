package validate

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func Create(obj *metric.CreateI) error {
	if len(obj.Datapoint) != 2 {
		return tracer.Maskf(invalidInputError, "createi.datapoint must have two elements")
	}

	if obj.UpdateId == "" {
		return tracer.Maskf(invalidInputError, "createi.update_id must not be empty")
	}

	return nil
}
