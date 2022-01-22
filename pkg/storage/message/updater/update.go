package updater

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apicommon/pkg/schema"
	"github.com/venturemark/apigengo/pkg/pbf/message"
	"github.com/xh3b4sd/tracer"
)

func (u *Updater) Update(req *message.UpdateI) (*message.UpdateO, error) {
	var err error

	var mek *key.Key
	{
		mek = key.Message(req.Obj[0].Metadata)
	}

	var cur []byte
	{
		k := mek.List()
		s := mek.ID().F()

		str, err := u.redigo.Sorted().Search().Score(k, s, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		mes := &schema.Message{}
		err = json.Unmarshal([]byte(str[0]), mes)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cur, err = json.Marshal(mes)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var pat []byte
	{
		var p []map[string]string

		for _, j := range req.Obj[0].Jsnpatch {
			m := map[string]string{
				"op":    j.GetOpe(),
				"path":  j.GetPat(),
				"value": j.GetVal(),
			}

			p = append(p, m)
		}

		pat, err = json.Marshal(p)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var val string
	{
		patch, err := jsonpatch.DecodePatch(pat)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		des, err := patch.Apply(cur)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = string(des)
	}

	var upd bool
	{
		k := mek.List()
		v := val
		s := mek.ID().F()

		upd, err = u.redigo.Sorted().Update().Value(k, v, s)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *message.UpdateO
	{
		res = &message.UpdateO{
			Obj: []*message.UpdateO_Obj{
				{
					Metadata: map[string]string{
						metadata.MessageID: mek.ID().S(),
					},
				},
			},
		}

		if upd {
			res.Obj[0].Metadata[metadata.MessageStatus] = "updated"
		}
	}

	return res, nil
}
