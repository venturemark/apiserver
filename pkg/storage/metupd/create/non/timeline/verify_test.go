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
		// Case 0 ensures that create input without filter is not valid.
		{
			obj: &metupd.CreateI{},
		},
		// Case 1 ensures that create input with an empty filter object is not
		// valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{},
			},
		},
		// Case 2 ensures that create input with empty operator and property
		// objects is not valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{},
					Property: []*metupd.SearchI_Filter_Property{},
				},
			},
		},
		// Case 3 ensures that create input with timestamp properties is not
		// valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timestamp: "1605025038",
						},
					},
				},
			},
		},
		// Case 4 ensures that create input with timestamp properties is not
		// valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timestamp: "1605025038",
						},
						{
							Timeline: "tml-al9qy",
						},
					},
				},
			},
		},
		// Case 5 ensures that create input with timestamp properties is not
		// valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timeline:  "tml-al9qy",
							Timestamp: "1605025038",
						},
						{
							Timeline: "tml-al9qy",
						},
					},
				},
			},
		},
		// Case 6 ensures that create input with multiple operators is not
		// valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timeline: "tml-al9qy",
						},
						{
							Timeline: "tml-i45kj",
						},
					},
				},
			},
		},
		// Case 7 ensures that create input with duplicated properties is not
		// valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timeline: "tml-al9qy",
						},
						{
							Timeline: "tml-al9qy",
						},
						{
							Timeline: "tml-al9qy",
						},
					},
				},
			},
		},
		// Case 8 ensures that create input with a single property is not valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timeline: "tml-al9qy",
						},
					},
				},
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
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timeline: "tml-kj3h4",
						},
					},
				},
			},
		},
		// Case 1 ensures that create input can be considered valid.
		{
			obj: &metupd.CreateI{
				Filter: &metupd.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metupd.SearchI_Filter_Property{
						{
							Timeline: "tml-al9qy",
						},
						{
							Timeline: "tml-i45kj",
						},
						{
							Timeline: "tml-kj3h4",
						},
					},
				},
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
