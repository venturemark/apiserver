package update

import (
	"github.com/venturemark/apigengo/pkg/pbf/metric"
)

func (u *Update) Verify(obj *metric.SearchI) bool {
	return false
}
