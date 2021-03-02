package updater

import (
	"github.com/venturemark/apigengo/pkg/pbf/role"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Verify(req *role.UpdateI) (bool, error) {
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
