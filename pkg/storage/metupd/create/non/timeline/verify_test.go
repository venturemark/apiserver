package timeline

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	loggerfake "github.com/xh3b4sd/logger/fake"
	redigofake "github.com/xh3b4sd/redigo/fake"
)

func Test_Timeline_Verify_Invalid(t *testing.T) {
	testCases := []struct {
		obj *metupd.CreateI
	}{
		// Case 0 ensures that create input without any information provided is
		// not valid.
		{
			obj: &metupd.CreateI{},
		},
		// Case 1 ensures that create input with only a single datapoint is not
		// valid.
		{
			obj: &metupd.CreateI{
				Datapoint: []int64{
					1604959525,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
		},
		// Case 2 ensures that create input with too many datapoints is not
		// valid.
		{
			obj: &metupd.CreateI{
				Datapoint: []int64{
					1604959525,
					85,
					2387,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
		},
		// Case 3 ensures that create input without text is not valid.
		{
			obj: &metupd.CreateI{
				Datapoint: []int64{
					1604959525,
					85,
				},
				Timeline: "tml-al9qy",
			},
		},
		// Case 4 ensures that create input without timeline is not valid.
		{
			obj: &metupd.CreateI{
				Datapoint: []int64{
					1604959525,
					85,
				},
				Text: "Lorem ipsum ...",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var tml *Timeline
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: redigofake.New(),
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok := tml.Verify(tc.obj)

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Timeline_Verify_Valid(t *testing.T) {
	testCases := []struct {
		obj *metupd.CreateI
	}{
		// Case 0 ensures that create input can be considered valid.
		{
			obj: &metupd.CreateI{
				Datapoint: []int64{
					1605025038,
					85,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
		},
		// Case 0 ensures that create input can be considered valid.
		{
			obj: &metupd.CreateI{
				Datapoint: []int64{
					1604959525,
					2500,
				},
				Text:     "foo bar #hashtag",
				Timeline: "tml-i45kj",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var tml *Timeline
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: redigofake.New(),
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok := tml.Verify(tc.obj)

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
