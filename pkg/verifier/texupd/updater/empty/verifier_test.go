package empty

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apicommon/pkg/metadata"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
)

func Test_Empty_Verify_False(t *testing.T) {
	testCases := []struct {
		req *texupd.UpdateI
	}{
		// Case 0 ensures that empty update input is not valid.
		{
			req: &texupd.UpdateI{},
		},
		// Case 1 ensures that empty update input is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{},
			},
		},
		// Case 2 ensures that empty update input is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Property: &texupd.UpdateI_Obj_Property{},
				},
			},
		},
		// Case 3 ensures that update input without metadata is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 4 ensures that update input without venture ID in the metadata
		// is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "<id>",
						metadata.UpdateID:   "<id>",
					},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 5 ensures that update input without timeline ID is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.VentureID: "org-al9qy",
					},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 7 ensures that update input with empty text is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.VentureID:  "org-al9qy",
						metadata.TimelineID: "1606329189",
					},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "",
					},
				},
			},
		},
		// Case 8 ensures that update input without update ID in the
		// metadata is not valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.VentureID:  "<id>",
						metadata.TimelineID: "<id>",
					},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
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
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Empty_Verify_True(t *testing.T) {
	testCases := []struct {
		req *texupd.UpdateI
	}{
		// Case 0 ensures that update input with text is valid.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.VentureID:  "<id>",
						metadata.TimelineID: "<id>",
						metadata.UpdateID:   "<id>",
					},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
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
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
