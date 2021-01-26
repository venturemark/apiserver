package element

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Split(t *testing.T) {
	testCases := []struct {
		mid float64
		oid string
		tex string
		rid string
		usr string
		bas string
	}{
		// Case 0 ensures the data format encoding can be split into its
		// original input.
		{
			mid: 1605559909,
			oid: "venturemark",
			tex: "foo bar",
			rid: "rid-al9qy",
			usr: "xh3b4sd",
			bas: "1605559909,dmVudHVyZW1hcms=,Zm9vIGJhcg==,cmlkLWFsOXF5,eGgzYjRzZA==",
		},
		// Case 1 ensures the data format encoding can be split into its
		// original input.
		{
			mid: 1605559912,
			oid: "venturemark",
			tex: "foo, bar",
			rid: "rid-pl3d7",
			usr: "xh3b4sd",
			bas: "1605559912,dmVudHVyZW1hcms=,Zm9vLCBiYXI=,cmlkLXBsM2Q3,eGgzYjRzZA==",
		},
		// Case 2 ensures the data format encoding can be split into its
		// original input.
		{
			mid: 1605858909,
			oid: "venturemark",
			tex: "foo, bar | baz ??? 2i376 kj ---..,23r2d3kj^^` boom boom",
			rid: "rid-al9qy",
			usr: "xh3b4sd",
			bas: "1605858909,dmVudHVyZW1hcms=,Zm9vLCBiYXIgfCBiYXogPz8/IDJpMzc2IGtqIC0tLS4uLDIzcjJkM2tqXl5gIGJvb20gYm9vbQ==,cmlkLWFsOXF5,eGgzYjRzZA==",
		},
		// Case 3 ensures the data format encoding can be split into its
		// original input.
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
			mid, oid, tex, rid, usr, err := Split(tc.bas)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.mid, mid) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.mid, mid))
			}
			if !reflect.DeepEqual(tc.oid, oid) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.oid, oid))
			}
			if !reflect.DeepEqual(tc.tex, tex) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.tex, tex))
			}
			if !reflect.DeepEqual(tc.rid, rid) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.rid, rid))
			}
			if !reflect.DeepEqual(tc.usr, usr) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.usr, usr))
			}
		})
	}
}
