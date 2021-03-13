package deleter

import (
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (c *Deleter) Delete(req *message.DeleteI) (*message.DeleteO, error) {
	var err error

	var mek *key.Key
	{
		mek = key.Message(req.Obj[0].Metadata)
	}

	{
		k := mek.List()
		s := mek.ID().F()

		err = c.redigo.Sorted().Delete().Score(k, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.DeleteO
	{
		res = &message.DeleteO{
			Obj: []*message.DeleteO_Obj{
				{
					Metadata: map[string]string{
						metadata.MessageStatus: "deleted",
					},
				},
			},
		}
	}

	return res, nil
}
