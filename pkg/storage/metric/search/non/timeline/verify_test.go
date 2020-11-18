package timeline

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	"github.com/venturemark/apiserver/pkg/metadata"
	loggerfake "github.com/xh3b4sd/logger/fake"
	redigofake "github.com/xh3b4sd/redigo/fake"
)

func Test_Timeline_Verify_Bool_False(t *testing.T) {
	testCases := []struct {
		req *metric.SearchI
	}{
		// Case 0 ensures that search input without object list is not valid.
		{
			req: &metric.SearchI{},
		},
		// Case 1 ensures that search input with an empty object list is not
		// valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{},
			},
		},
		// Case 2 ensures that search input with property is not valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Property: &metric.SearchI_Obj_Property{},
					},
				},
			},
		},
		// Case 3 ensures that search input with property is not valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Property: &metric.SearchI_Obj_Property{},
					},
					{
						Property: &metric.SearchI_Obj_Property{},
					},
					{
						Property: &metric.SearchI_Obj_Property{},
					},
				},
			},
		},
		// Case 4 ensures that search input with timestamp is not valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Unixtime: "1605025038",
						},
					},
				},
			},
		},
		// Case 5 ensures that search input with timestamp is not valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Unixtime: "1605025038",
						},
					},
					{
						Metadata: map[string]string{
							metadata.Unixtime: "1605025038",
						},
					},
					{
						Metadata: map[string]string{
							metadata.Unixtime: "1605025038",
						},
					},
				},
			},
		},
		// Case 6 ensures that search input with multiple timestamps is not
		// valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
					{
						Metadata: map[string]string{
							metadata.Unixtime: "1605025038",
						},
					},
					{
						Metadata: map[string]string{
							metadata.Unixtime: "1605025038",
						},
					},
				},
			},
		},
		// Case 7 ensures that search input with multiple timelines is not valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
				},
			},
		},
		// Case 8 ensures that search input with timestamp and timeline is not
		// valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
							metadata.Unixtime: "1605025038",
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

			ok, err := tml.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Timeline_Verify_Bool_True(t *testing.T) {
	testCases := []struct {
		req *metric.SearchI
	}{
		// Case 0 ensures that search input with timeline is valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
				},
			},
		},
		// Case 1 ensures that search input with timeline is valid.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-i45kj",
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

			ok, err := tml.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
