package updater

import (
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Verify(req *texupd.UpdateI) (bool, error) {
	for _, v := range u.verify {
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
