package length

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

func Test_Length_Verify_False(t *testing.T) {
	testCases := []struct {
		req *timeline.CreateI
	}{
		// Case 0 ensures that create input with too long of a description is
		// not valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Desc: strings.Repeat("0123456789", 29),
					},
				},
			},
		},
		// Case 1 ensures that create input with too long of a description is
		// not valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Desc: strings.Repeat("0123456789", 40),
					},
				},
			},
		},
		// Case 2 ensures that create input with too long of a description is
		// not valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Desc: strings.Repeat("0123456789", 100),
					},
				},
			},
		},
		// Case 3 ensures that create input with too long of a name is not
		// valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Name: strings.Repeat("0123456789", 4),
					},
				},
			},
		},
		// Case 4 ensures that create input with too long of a name is not
		// valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Name: strings.Repeat("0123456789", 40),
					},
				},
			},
		},
		// Case 5 ensures that create input with too long of a name is not
		// valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Name: strings.Repeat("0123456789", 100),
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

func Test_Length_Verify_True(t *testing.T) {
	testCases := []struct {
		req *timeline.CreateI
	}{
		// Case 0 ensures that empty create input is valid.
		{
			req: &timeline.CreateI{},
		},
		// Case 1 ensures that empty create input is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{},
			},
		},
		// Case 2 ensures that empty create input is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{},
				},
			},
		},
		// Case 3 ensures that create input with description is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Desc: strings.Repeat("0123456789", 1),
					},
				},
			},
		},
		// Case 4 ensures that create input with description is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Desc: strings.Repeat("0123456789", 10),
					},
				},
			},
		},
		// Case 5 ensures that create input with description is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Desc: strings.Repeat("0123456789", 28),
					},
				},
			},
		},
		// Case 6 ensures that create input with name is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Name: strings.Repeat("0123456789", 1),
					},
				},
			},
		},
		// Case 7 ensures that create input with name is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Name: strings.Repeat("0123456789", 2),
					},
				},
			},
		},
		// Case 8 ensures that create input with name is valid.
		{
			req: &timeline.CreateI{
				Obj: &timeline.CreateI_Obj{
					Property: &timeline.CreateI_Obj_Property{
						Name: strings.Repeat("0123456789", 3),
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
