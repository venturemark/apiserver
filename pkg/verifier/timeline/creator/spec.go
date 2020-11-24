package creator

import "github.com/venturemark/apigengo/pkg/pbf/timeline"

type Interface interface {
	Verify(req *timeline.CreateI) (bool, error)
}
