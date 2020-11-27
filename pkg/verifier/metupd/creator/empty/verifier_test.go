package empty

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Empty_Verify_False(t *testing.T) {
	testCases := []struct {
		req *metupd.CreateI
	}{
		// Case 0 ensures that empty update input is not valid.
		{
			req: &metupd.CreateI{},
		},
		// Case 1 ensures that empty update input is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{},
			},
		},
		// Case 2 ensures that empty update input is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Property: &metupd.CreateI_Obj_Property{},
				},
			},
		},
		// Case 3 ensures that create input without metadata is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 4 ensures that create input without user ID in the metadata
		// is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 5 ensures that create input without timeline ID is not
		// valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.UserID: "usr-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 6 ensures that update input with empty data is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{},
					},
				},
			},
		},
		// Case 7 ensures that update input with empty data and empty text is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{},
						Text: "",
					},
				},
			},
		},
		// Case 8 ensures that update input with empty data is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{},
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
		req *metupd.CreateI
	}{
		// Case 0 ensures that update input with data and text is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 1 ensures that update input with data and text is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1605559909",
						metadata.UserID:     "usr-w4ndz",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{},
							{},
							{},
						},
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
