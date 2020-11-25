package updater

import "github.com/venturemark/apigengo/pkg/pbf/timeline"

type Interface interface {
	Verify(req *timeline.UpdateI) (bool, error)
}
