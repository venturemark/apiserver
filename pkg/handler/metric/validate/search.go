package validate

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/xh3b4sd/tracer"
)

func Search(obj *metric.SearchI) error {
	return tracer.Maskf(invalidInputError, "")
}
