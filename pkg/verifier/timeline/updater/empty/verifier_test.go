package empty

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Empty_Verify_False(t *testing.T) {
	testCases := []struct {
		req *timeline.UpdateI
	}{
		// Case 0 ensures that create input without metadata is not valid.
		{
			req: &timeline.UpdateI{},
		},
		// Case 1 ensures that create input without metadata is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{},
			},
		},
		// Case 2 ensures that create input without metadata is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{},
			},
		},
		// Case 3 ensures that create input without metadata is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{},
				},
			},
		},
		// Case 4 ensures that create input without user ID in the metadata
		// is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{
						"foo": "bar",
					},
				},
			},
		},
		// Case 5 ensures that create input without user ID is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
					},
					Property: &timeline.UpdateI_Obj_Property{
						Name: "foo",
					},
				},
			},
		},
		// Case 6 ensures that create input without timeline ID is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
					},
					Property: &timeline.UpdateI_Obj_Property{
						Name: "foo",
					},
				},
			},
		},
		// Case 7 ensures that create input without object properties is not
		// valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "1606329189",
					},
					Property: &timeline.UpdateI_Obj_Property{},
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
		req *timeline.UpdateI
	}{
		// Case 0 ensures that create input with user ID is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "1606329189",
					},
					Property: &timeline.UpdateI_Obj_Property{
						Name: "mmr",
					},
				},
			},
		},
		// Case 1 ensures that create input with user ID is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "1605559909",
					},
					Property: &timeline.UpdateI_Obj_Property{
						Name: "MMR",
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
