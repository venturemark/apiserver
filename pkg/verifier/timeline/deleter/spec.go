package deleter

import "github.com/venturemark/apigengo/pkg/pbf/timeline"

type Interface interface {
	Verify(req *timeline.DeleteI) (bool, error)
}
