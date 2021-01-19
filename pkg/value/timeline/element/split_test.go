package element

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Element_Split(t *testing.T) {
	testCases := []struct {
		uni float64
		des string
		nam string
		sta string
		bas string
	}{
		// Case 0 ensures the data format encoding can be split into its
		// original input.
		{
			uni: 1605559909,
			des: "foo bar",
			nam: "mrr",
			sta: "archived",
			bas: "1605559909,Zm9vIGJhcg==,bXJy,YXJjaGl2ZWQ=",
		},
		// Case 1 ensures the data format encoding can be split into its
		// original input.
		{
			uni: 1605559912,
			des: "foo, bar",
			nam: "rid-pl3d7",
			sta: "",
			bas: "1605559912,Zm9vLCBiYXI=,cmlkLXBsM2Q3,",
		},
		// Case 2 ensures the data format encoding can be split into its
		// original input.
		{
			uni: 1605858909,
			des: "",
			nam: "rid-al9qy",
			sta: "rid-al9qy",
			bas: "1605858909,,cmlkLWFsOXF5,cmlkLWFsOXF5",
		},
		// Case 3 ensures the data format encoding can be split into its
		// original input.
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
			uni, des, nam, sta, err := Split(tc.bas)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.uni, uni) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.uni, uni))
			}
			if !reflect.DeepEqual(tc.des, des) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.des, des))
			}
			if !reflect.DeepEqual(tc.nam, nam) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.nam, nam))
			}
			if !reflect.DeepEqual(tc.sta, sta) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.sta, sta))
			}
		})
	}
}
