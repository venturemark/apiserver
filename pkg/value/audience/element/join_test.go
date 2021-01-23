package element

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Join(t *testing.T) {
	testCases := []struct {
		aid float64
		nam string
		tim []string
		usr []string
		bas string
	}{
		// Case 0 ensures a simple string can be joined to the data format
		// encoding.
		{
			aid: 1605559909,
			nam: "foo bar",
			tim: []string{
				"tim-plf5j",
			},
			usr: []string{
				"usr-al9qy",
			},
			bas: "1605559909,Zm9vIGJhcg==,dGltLXBsZjVq,dXNyLWFsOXF5",
		},
		// Case 1 ensures a comma separarted string can be joined to the data
		// format encoding.
		{
			aid: 1605559912,
			nam: "foo, bar",
			tim: []string{
				"tim-plf5j",
			},
			usr: []string{
				"usr-al9qy",
				"usr-pl3d7",
			},
			bas: "1605559912,Zm9vLCBiYXI=,dGltLXBsZjVq,dXNyLWFsOXF5LHVzci1wbDNkNw==",
		},
		// Case 2 ensures a complex string can be joined to the data format
		// encoding.
		{
			aid: 1605858909,
			nam: "foo, bar | baz ??? 2i376 kj ---..,23r2d3kj^^` boom boom",
			tim: []string{
				"tim-plf5j",
			},
			usr: []string{},
			bas: "1605858909,Zm9vLCBiYXIgfCBiYXogPz8/IDJpMzc2IGtqIC0tLS4uLDIzcjJkM2tqXl5gIGJvb20gYm9vbQ==,dGltLXBsZjVq,",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bas := Join(tc.aid, tc.nam, tc.tim, tc.usr)

			if tc.bas != bas {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.bas, bas))
			}
		})
	}
}
