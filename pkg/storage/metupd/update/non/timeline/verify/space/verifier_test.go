package space

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

func Test_Space_Verify_False(t *testing.T) {
	testCases := []struct {
		req *metupd.UpdateI
	}{
		// Case 0 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{},
						},
					},
				},
			},
		},
		// Case 1 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{},
							{},
							{},
						},
					},
				},
			},
		},
		// Case 2 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "",
							},
						},
					},
				},
			},
		},
		// Case 3 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "",
							},
							{
								Space: "",
							},
							{
								Space: "",
							},
						},
					},
				},
			},
		},
		// Case 4 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "x",
							},
							{
								Space: "y",
							},
							{
								Space: "",
							},
						},
					},
				},
			},
		},
		// Case 5 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "y",
							},
							{
								Space: "",
							},
						},
					},
				},
			},
		},
		// Case 6 ensures that update input with the reserved dimensional space
		// t is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "t",
							},
						},
					},
				},
			},
		},
		// Case 7 ensures that update input with the reserved dimensional space
		// t is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "y",
							},
							{
								Space: "t",
							},
							{
								Space: "x",
							},
						},
					},
				},
			},
		},
		// Case 8 ensures that update input with the duplicated dimensional
		// spaces is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "x",
							},
							{
								Space: "x",
							},
						},
					},
				},
			},
		},
		// Case 9 ensures that update input with the duplicated dimensional
		// spaces is not valid. Note that this case is particularly tricky
		// because some implementations to check for duplication do not find
		// duplications if the first space, here y, is not part of the
		// duplication itself.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "y",
							},
							{
								Space: "z",
							},
							{
								Space: "z",
							},
						},
					},
				},
			},
		},
		// Case 10 ensures that update input with the dimensional spaces
		// identified with something else than single letters is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "9",
							},
							{
								Space: "x",
							},
						},
					},
				},
			},
		},
		// Case 11 ensures that update input with the dimensional spaces
		// identified with something else than single letters is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: " ",
							},
						},
					},
				},
			},
		},
		// Case 12 ensures that update input with the dimensional spaces
		// identified with something else than single letters is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "  f ",
							},
						},
					},
				},
			},
		},
		// Case 13 ensures that update input with the dimensional spaces
		// identified with something else than single letters is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "foo",
							},
						},
					},
				},
			},
		},
		// Case 14 ensures that update input with the reserved dimensional space
		// t is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "t",
							},
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

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Space_Verify_True(t *testing.T) {
	testCases := []struct {
		req *metupd.UpdateI
	}{
		// Case 0 ensures that empty update input is valid.
		{
			req: &metupd.UpdateI{},
		},
		// Case 1 ensures that empty update input is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{},
			},
		},
		// Case 2 ensures that update input with data is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "x",
							},
						},
					},
				},
			},
		},
		// Case 3 ensures that update input with data is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "x",
							},
							{
								Space: "y",
							},
							{
								Space: "z",
							},
						},
					},
				},
			},
		},
		// Case 4 ensures that update input with data is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "g",
							},
						},
					},
				},
			},
		},
		// Case 5 ensures that update input with data is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "g",
							},
							{
								Space: "p",
							},
							{
								Space: "v",
							},
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
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}
