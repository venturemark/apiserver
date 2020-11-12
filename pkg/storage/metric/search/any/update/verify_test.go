package update

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	loggerfake "github.com/xh3b4sd/logger/fake"
	redigofake "github.com/xh3b4sd/redigo/fake"
)

func Test_Update_Verify_Invalid(t *testing.T) {
	testCases := []struct {
		obj *metric.SearchI
	}{
		// Case 0 ensures that search input without filter is not valid.
		{
			obj: &metric.SearchI{},
		},
		// Case 1 ensures that search input with an empty filter object is not
		// valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{},
			},
		},
		// Case 2 ensures that search input with empty operator and property
		// objects is not valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{},
					Property: []*metric.SearchI_Filter_Property{},
				},
			},
		},
		// Case 3 ensures that search input with timestamp properties is not
		// valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							Timestamp: "1605025038",
						},
					},
				},
			},
		},
		// Case 4 ensures that search input with timestamp properties is not
		// valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							Timestamp: "1605025038",
						},
						{
							UpdateId: "upd-al9qy",
						},
					},
				},
			},
		},
		// Case 5 ensures that search input with timestamp properties is not
		// valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							Timestamp: "1605025038",
							UpdateId:  "upd-al9qy",
						},
						{
							UpdateId: "upd-al9qy",
						},
					},
				},
			},
		},
		// Case 6 ensures that search input with multiple operators is not
		// valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							UpdateId: "upd-al9qy",
						},
						{
							UpdateId: "upd-i45kj",
						},
					},
				},
			},
		},
		// Case 7 ensures that search input with duplicated properties is not
		// valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							UpdateId: "upd-al9qy",
						},
						{
							UpdateId: "upd-al9qy",
						},
						{
							UpdateId: "upd-al9qy",
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var u *Update
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: redigofake.New(),
				}

				u, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok := u.Verify(tc.obj)

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Update_Verify_Valid(t *testing.T) {
	testCases := []struct {
		obj *metric.SearchI
	}{
		// Case 0 ensures that search input can be considered valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							UpdateId: "upd-kj3h4",
						},
					},
				},
			},
		},
		// Case 1 ensures that search input can be considered valid.
		{
			obj: &metric.SearchI{
				Filter: &metric.SearchI_Filter{
					Operator: []string{
						"any",
					},
					Property: []*metric.SearchI_Filter_Property{
						{
							UpdateId: "upd-al9qy",
						},
						{
							UpdateId: "upd-i45kj",
						},
						{
							UpdateId: "upd-kj3h4",
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var u *Update
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: redigofake.New(),
				}

				u, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok := u.Verify(tc.obj)

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
