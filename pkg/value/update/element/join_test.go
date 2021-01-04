package element

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Join(t *testing.T) {
	testCases := []struct {
		tex string
		uni float64
		bas string
	}{
		// Case 0 ensures a simple string can be joined to the data format
		// encoding.
		{
			tex: "foo bar",
			uni: 1605559909,
			bas: "1605559909,Zm9vIGJhcg==",
		},
		// Case 1 ensures a comma separarted string can be joined to the data
		// format encoding.
		{
			tex: "foo, bar",
			uni: 1605559912,
			bas: "1605559912,Zm9vLCBiYXI=",
		},
		// Case 2 ensures a complex string can be joined to the data format
		// encoding.
		{
			tex: "foo, bar | baz ??? 2i376 kj ---..,23r2d3kj^^` boom boom",
			uni: 1605858909,
			bas: "1605858909,Zm9vLCBiYXIgfCBiYXogPz8/IDJpMzc2IGtqIC0tLS4uLDIzcjJkM2tqXl5gIGJvb20gYm9vbQ==",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bas := Join(tc.uni, tc.tex)

			if !reflect.DeepEqual(tc.bas, bas) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.bas, bas))
			}
		})
	}
}
