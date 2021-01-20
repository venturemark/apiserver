package state

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
)

func Test_Empty_Verify_False(t *testing.T) {
	testCases := []struct {
		req *timeline.UpdateI
	}{
		// Case 0 ensures that update input with unknown state is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Stat: toStringP("foo"),
					},
				},
			},
		},
		// Case 1 ensures that update input with unknown state is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Stat: toStringP("bar"),
					},
				},
			},
		},
		// Case 2 ensures that update input with unknown state is not valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Stat: toStringP("archive"),
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
		req *timeline.UpdateI
	}{
		// Case 0 ensures that update input with known state is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Stat: toStringP(""),
					},
				},
			},
		},
		// Case 1 ensures that update input with known state is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Stat: toStringP("active"),
					},
				},
			},
		},
		// Case 2 ensures that update input with known state is valid.
		{
			req: &timeline.UpdateI{
				Obj: &timeline.UpdateI_Obj{
					Property: &timeline.UpdateI_Obj_Property{
						Stat: toStringP("archived"),
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

func toStringP(s string) *string {
	return &s
}
