package element

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Join(t *testing.T) {
	testCases := []struct {
		uid float64
		oid string
		tex string
		usr string
		bas string
	}{
		// Case 0 ensures a simple string can be joined to the data format
		// encoding.
		{
			uid: 1605559909,
			oid: "venturemark",
			tex: "foo bar",
			usr: "xh3b4sd",
			bas: "1605559909,dmVudHVyZW1hcms=,Zm9vIGJhcg==,eGgzYjRzZA==",
		},
		// Case 1 ensures a comma separarted string can be joined to the data
		// format encoding.
		{
			uid: 1605559912,
			oid: "venturemark",
			tex: "foo, bar",
			usr: "xh3b4sd",
			bas: "1605559912,dmVudHVyZW1hcms=,Zm9vLCBiYXI=,eGgzYjRzZA==",
		},
		// Case 2 ensures a complex string can be joined to the data format
		// encoding.
		{
			uid: 1605858909,
			oid: "venturemark",
			tex: "foo, bar | baz ??? 2i376 kj ---..,23r2d3kj^^` boom boom",
			usr: "xh3b4sd",
			bas: "1605858909,dmVudHVyZW1hcms=,Zm9vLCBiYXIgfCBiYXogPz8/IDJpMzc2IGtqIC0tLS4uLDIzcjJkM2tqXl5gIGJvb20gYm9vbQ==,eGgzYjRzZA==",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bas := Join(tc.uid, tc.oid, tc.tex, tc.usr)

			if !reflect.DeepEqual(tc.bas, bas) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.bas, bas))
			}
		})
	}
}
