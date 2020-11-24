package element

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_Data_Join tests parsing of data elements of a sorted set associated with
// a timeline. Note that these tests do also verify the validity of floating
// point numbers as well as the ordering by dimensional space while the order of
// provided datapoints remains intact.
//
//     go test ./pkg/value/metric/timeline/data -run Test_Data_Join -update
//
func Test_Data_Join(t *testing.T) {
	testCases := []struct {
		str string
		uni float64
		val []Interface
	}{
		// Case 0 ensures that data elements of a sorted set associated with a
		// timeline are parsed properly.
		{
			str: "1605559909:y,23",
			uni: 1605559909,
			val: []Interface{
				Wrapper{
					Space: "y",
					Value: []float64{
						23,
					},
				},
			},
		},
		// Case 1 ensures that data elements of a sorted set associated with a
		// timeline are parsed properly. Note that this case provides an
		// unordered list where the resulting string is ordered by dimensional
		// space while the provided datapoints stay untouched.
		{
			str: "1605559909:x,0.5,0.7,600:y,23,8.553,300:z,15.3,1,9040",
			uni: 1605559909,
			val: []Interface{
				Wrapper{
					Space: "y",
					Value: []float64{
						23,
						8.553,
						300,
					},
				},
				Wrapper{
					Space: "z",
					Value: []float64{
						15.3,
						1,
						9040,
					},
				},
				Wrapper{
					Space: "x",
					Value: []float64{
						0.5,
						0.7,
						600,
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			str := Join(tc.uni, tc.val)

			if !reflect.DeepEqual(tc.str, str) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.str, str))
			}
		})
	}
}
