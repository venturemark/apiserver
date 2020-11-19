package data

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update", false, "update .golden files")

// Test_Data_Split tests parsing of data elements of a sorted set associated
// with a timeline. Note that these tests do also verify the validity of
// floating point numbers as well as the injection of the reserved dimensional
// space t. Here t is time, tracking the unix timestamp for each datapoint.
//
//     go test ./pkg/value/metric/timeline/data -run Test_Data_Split -update
//
func Test_Data_Split(t *testing.T) {
	testCases := []struct {
		str string
	}{
		// Case 0 ensures that data elements of a sorted set associated with a
		// timeline are parsed properly.
		{
			str: "1605559909:y,23",
		},
		// Case 1 ensures that data elements of a sorted set associated with a
		// timeline are parsed properly. Note that this case uses floating point
		// numbers in multiple dimensions.
		{
			str: "1605559909:y,23,8.553,300:x,0.5,0.7,600",
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var actual []byte
			{
				uni, val, err := Split(tc.str)
				if err != nil {
					t.Fatal(err)
				}

				b, err := json.MarshalIndent(val, "", "    ")
				if err != nil {
					t.Fatal(err)
				}

				b = append(b, '\n', '\n')
				b = strconv.AppendFloat(b, uni, 'f', -1, 64)
				b = append(b, '\n')

				actual = b
			}

			p := filepath.Join("testdata/split", fileName(i))
			if *update {
				err := ioutil.WriteFile(p, []byte(actual), 0600)
				if err != nil {
					t.Fatal(err)
				}
			}

			expected, err := ioutil.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(expected, []byte(actual)) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(actual), string(expected)))
			}
		})
	}
}

func fileName(i int) string {
	return "case-" + strconv.Itoa(i) + ".golden"
}
