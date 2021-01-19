package element

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Join(t *testing.T) {
	testCases := []struct {
		uni float64
		des string
		nam string
		sta string
		bas string
	}{
		// Case 0 ensures a simple string can be joined to the data format
		// encoding.
		{
			uni: 1605559909,
			des: "foo bar",
			nam: "mrr",
			sta: "archived",
			bas: "1605559909,Zm9vIGJhcg==,bXJy,YXJjaGl2ZWQ=",
		},
		// Case 1 ensures a comma separarted string can be joined to the data
		// format encoding.
		{
			uni: 1605559912,
			des: "foo, bar",
			nam: "rid-pl3d7",
			sta: "",
			bas: "1605559912,Zm9vLCBiYXI=,cmlkLXBsM2Q3,",
		},
		// Case 2 ensures a complex string can be joined to the data format
		// encoding.
		{
			uni: 1605858909,
			des: "",
			nam: "rid-al9qy",
			sta: "rid-al9qy",
			bas: "1605858909,,cmlkLWFsOXF5,cmlkLWFsOXF5",
		},
		// Case 3 ensures a simple string can be joined to the data format
		// encoding.
		{
			uni: 1605559909,
			des: "",
			nam: "foo",
			sta: "",
			bas: "1605559909,,Zm9v,",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bas := Join(tc.uni, tc.des, tc.nam, tc.sta)

			if tc.bas != bas {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.bas, bas))
			}
		})
	}
}
