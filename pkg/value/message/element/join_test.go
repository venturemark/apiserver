package element

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Join(t *testing.T) {
	testCases := []struct {
		uni float64
		tex string
		rid string
		bas string
	}{
		// Case 0 ensures a simple string can be joined to the data format
		// encoding.
		{
			uni: 1605559909,
			tex: "foo bar",
			rid: "rid-al9qy",
			bas: "1605559909,Zm9vIGJhcg==,cmlkLWFsOXF5",
		},
		// Case 1 ensures a comma separarted string can be joined to the data
		// format encoding.
		{
			uni: 1605559912,
			tex: "foo, bar",
			rid: "rid-pl3d7",
			bas: "1605559912,Zm9vLCBiYXI=,cmlkLXBsM2Q3",
		},
		// Case 2 ensures a complex string can be joined to the data format
		// encoding.
		{
			uni: 1605858909,
			tex: "foo, bar | baz ??? 2i376 kj ---..,23r2d3kj^^` boom boom",
			rid: "rid-al9qy",
			bas: "1605858909,Zm9vLCBiYXIgfCBiYXogPz8/IDJpMzc2IGtqIC0tLS4uLDIzcjJkM2tqXl5gIGJvb20gYm9vbQ==,cmlkLWFsOXF5",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bas := Join(tc.uni, tc.tex, tc.rid)

			if tc.bas != bas {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.bas, bas))
			}
		})
	}
}
