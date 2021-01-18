package deleter

import (
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (c *Deleter) Verify(req *message.DeleteI) (bool, error) {
	for _, v := range c.verify {
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
