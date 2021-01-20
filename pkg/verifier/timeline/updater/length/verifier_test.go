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
		req *timeline.UpdateI
	}{
		// Case 0 ensures that update input with too long of a description is
		// not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Desc: toStringP(strings.Repeat("0123456789", 29)),
					},
				},
			},
		},
		// Case 1 ensures that update input with too long of a description is
		// not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Desc: toStringP(strings.Repeat("0123456789", 40)),
					},
				},
			},
		},
		// Case 2 ensures that update input with too long of a description is
		// not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Desc: toStringP(strings.Repeat("0123456789", 100)),
					},
				},
			},
		},
		// Case 3 ensures that update input with too long of a name is not
		// valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Name: toStringP(strings.Repeat("0123456789", 4)),
					},
				},
			},
		},
		// Case 4 ensures that update input with too long of a name is not
		// valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Name: toStringP(strings.Repeat("0123456789", 40)),
					},
				},
			},
		},
		// Case 5 ensures that update input with too long of a name is not
		// valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Name: toStringP(strings.Repeat("0123456789", 100)),
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
		req *timeline.UpdateI
	}{
		// Case 0 ensures that empty update input is valid.
		{
			req: &timeline.UpdateI{},
		},
		// Case 1 ensures that empty update input is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{},
			},
		},
		// Case 2 ensures that empty update input is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{},
				},
			},
		},
		// Case 3 ensures that update input with description is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Desc: toStringP(strings.Repeat("0123456789", 1)),
					},
				},
			},
		},
		// Case 4 ensures that update input with description is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Desc: toStringP(strings.Repeat("0123456789", 10)),
					},
				},
			},
		},
		// Case 5 ensures that update input with description is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Desc: toStringP(strings.Repeat("0123456789", 28)),
					},
				},
			},
		},
		// Case 6 ensures that update input with name is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Name: toStringP(strings.Repeat("0123456789", 1)),
					},
				},
			},
		},
		// Case 7 ensures that update input with name is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Name: toStringP(strings.Repeat("0123456789", 2)),
					},
				},
			},
		},
		// Case 8 ensures that update input with name is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Name: toStringP(strings.Repeat("0123456789", 3)),
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

func toStringP(s string) *string {
	return &s
}
