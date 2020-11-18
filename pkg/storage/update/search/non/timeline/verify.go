package timeline

import (
	"github.com/venturemark/apigengo/pkg/pbf/update"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func (t *Timeline) Verify(req *update.SearchI) (bool, error) {
	{
		// Any search request with api specifics is not valid at this point. We
		// will extend functionality here later.
		if req.Api != nil {
			return false, nil
		}
	}

	{
		// We need a single object with a single metadata label for the user's
		// timeline in order to fullfil the search request. We will extend
		// functionality here later.
		if len(req.Obj) != 1 {
			return false, nil
		}
		if len(req.Obj[0].Metadata) != 1 {
			return false, nil
		}
		if req.Obj[0].Metadata[metadata.Timeline] == "" {
			return false, nil
		}
	}

	{
		// Any search request with object property specifics is not valid at
		// this point. We will extend functionality here later.
		if req.Obj[0].Property != nil {
			return false, nil
		}
	}

	return true, nil
}
