package text

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
)

func Test_Text_Verify_False(t *testing.T) {
	testCases := []struct {
		req *texupd.CreateI
	}{
		// Case 0 ensures that update input with too long of a text is not
		// valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{
						Text: strings.Repeat("0123456789", 29),
					},
				},
			},
		},
		// Case 1 ensures that update input with too long of a text is not
		// valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{
						Text: strings.Repeat("0123456789", 40),
					},
				},
			},
		},
		// Case 2 ensures that update input with too long of a text is not
		// valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{
						Text: strings.Repeat("0123456789", 100),
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

func Test_Text_Verify_True(t *testing.T) {
	testCases := []struct {
		req *texupd.CreateI
	}{
		// Case 0 ensures that empty update input is valid.
		{
			req: &texupd.CreateI{},
		},
		// Case 1 ensures that empty update input is valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{},
			},
		},
		// Case 2 ensures that empty update input is valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{},
				},
			},
		},
		// Case 3 ensures that update input with text is valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{
						Text: strings.Repeat("0123456789", 1),
					},
				},
			},
		},
		// Case 4 ensures that update input with text is valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{
						Text: strings.Repeat("0123456789", 10),
					},
				},
			},
		},
		// Case 5 ensures that update input with text is valid.
		{
			req: &texupd.CreateI{
				Obj: &texupd.CreateI_Obj{
					Property: &texupd.CreateI_Obj_Property{
						Text: strings.Repeat("0123456789", 28),
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
