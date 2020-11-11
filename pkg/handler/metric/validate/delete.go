package validate

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func Delete(obj *metric.DeleteI) error {
	return tracer.Maskf(invalidInputError, "")
}
