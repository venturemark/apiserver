package validate

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func Update(obj *metric.UpdateI) error {
	return tracer.Maskf(invalidInputError, "")
}
