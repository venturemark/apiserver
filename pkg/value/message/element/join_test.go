package element

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Join(t *testing.T) {
	testCases := []struct {
		mid float64
		oid string
		tex string
		rid string
		usr string
		bas string
	}{
		// Case 0 ensures a simple string can be joined to the data format
		// encoding.
		{
			mid: 1605559909,
			oid: "venturemark",
			tex: "foo bar",
			rid: "rid-al9qy",
			usr: "xh3b4sd",
			bas: "1605559909,dmVudHVyZW1hcms=,Zm9vIGJhcg==,cmlkLWFsOXF5,eGgzYjRzZA==",
		},
		// Case 1 ensures a comma separarted string can be joined to the data
		// format encoding.
		{
			mid: 1605559912,
			oid: "venturemark",
			tex: "foo, bar",
			rid: "rid-pl3d7",
			usr: "xh3b4sd",
			bas: "1605559912,dmVudHVyZW1hcms=,Zm9vLCBiYXI=,cmlkLXBsM2Q3,eGgzYjRzZA==",
		},
		// Case 2 ensures a complex string can be joined to the data format
		// encoding.
		{
			mid: 1605858909,
			oid: "venturemark",
			tex: "foo, bar | baz ??? 2i376 kj ---..,23r2d3kj^^` boom boom",
			rid: "rid-al9qy",
			usr: "xh3b4sd",
			bas: "1605858909,dmVudHVyZW1hcms=,Zm9vLCBiYXIgfCBiYXogPz8/IDJpMzc2IGtqIC0tLS4uLDIzcjJkM2tqXl5gIGJvb20gYm9vbQ==,cmlkLWFsOXF5,eGgzYjRzZA==",
		},
		// Case 3 ensures a simple string can be joined to the data format
		// encoding.
		{
			mid: 1605559909,
			oid: "venturemark",
			tex: "foo bar",
			rid: "",
			usr: "xh3b4sd",
			bas: "1605559909,dmVudHVyZW1hcms=,Zm9vIGJhcg==,,eGgzYjRzZA==",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bas := Join(tc.mid, tc.oid, tc.tex, tc.rid, tc.usr)

			if tc.bas != bas {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.bas, bas))
			}
		})
	}
}
