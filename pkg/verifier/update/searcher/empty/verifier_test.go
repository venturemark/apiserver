package empty

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/update"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Empty_Verify_False(t *testing.T) {
	testCases := []struct {
		req *update.SearchI
	}{
		// Case 0 ensures that search input without metadata is not valid.
		{
			req: &update.SearchI{},
		},
		// Case 1 ensures that search input without metadata is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{},
			},
		},
		// Case 2 ensures that search input without metadata is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{},
				},
			},
		},
		// Case 3 ensures that search input without metadata is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{},
					},
				},
			},
		},
		// Case 4 ensures that search input without timeline ID in the metadata
		// is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							"foo": "bar",
						},
					},
				},
			},
		},
		// Case 5 ensures that search input with multiple objects is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.AudienceID: "aud-w4ndz",
							metadata.TimelineID: "1606329189",
						},
					},
					{
						Metadata: map[string]string{
							"foo": "bar",
						},
					},
				},
			},
		},
		// Case 6 ensures that search input with object properties is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.AudienceID: "aud-w4ndz",
							metadata.TimelineID: "1606329189",
						},
						Property: &update.SearchI_Obj_Property{},
					},
				},
			},
		},
		// Case 7 ensures that search input without timeline ID is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.AudienceID: "aud-w4ndz",
						},
						Property: &update.SearchI_Obj_Property{},
					},
				},
			},
		},
		// Case 8 ensures that search input without audience ID is not valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.TimelineID: "1606329189",
						},
						Property: &update.SearchI_Obj_Property{},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{}

				v, err = NewVerifier(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok, err := v.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(false, ok))
			}
		})
	}
}

func Test_Empty_Verify_True(t *testing.T) {
	testCases := []struct {
		req *update.SearchI
	}{
		// Case 0 ensures that search input with timeline ID is valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.AudienceID: "aud-w4ndz",
							metadata.TimelineID: "1606329189",
						},
					},
				},
			},
		},
		// Case 1 ensures that search input with timeline ID is valid.
		{
			req: &update.SearchI{
				Obj: []*update.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.AudienceID: "aud-al9qy",
							metadata.TimelineID: "1605559909",
						},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{}

				v, err = NewVerifier(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok, err := v.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(true, ok))
			}
		})
	}
}
