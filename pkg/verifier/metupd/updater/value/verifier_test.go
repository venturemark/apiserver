package value

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
)

func Test_Value_Verify_False(t *testing.T) {
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
								Value: []float64{},
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
								Value: []float64{},
							},
							{
								Value: []float64{},
							},
							{
								Value: []float64{},
							},
						},
					},
				},
			},
		},
		// Case 4 ensures that update input with inconsistent data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Value: []float64{
									83,
								},
							},
							{
								Value: []float64{
									15,
									8.3,
								},
							},
							{
								Value: []float64{
									2000,
									33.4,
									83,
								},
							},
						},
					},
				},
			},
		},
		// Case 5 ensures that update input with inconsistent data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Value: []float64{
									83,
								},
							},
							{
								Value: []float64{
									15,
									8.3,
								},
							},
							{
								Value: []float64{
									33.4,
								},
							},
						},
					},
				},
			},
		},
		// Case 6 ensures that update input with inconsistent data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Value: []float64{
									83,
								},
							},
							{
								Value: []float64{
									15,
									8.3,
								},
							},
							{
								Value: []float64{
									33.4,
									44,
								},
							},
						},
					},
				},
			},
		},
		// Case 7 ensures that update input with inconsistent data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Value: []float64{
									83,
									15,
								},
							},
							{
								Value: []float64{
									8.3,
									33.4,
								},
							},
							{
								Value: []float64{
									44,
								},
							},
						},
					},
				},
			},
		},
		// Case 8 ensures that update input with inconsistent data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Value: []float64{
									83,
									15,
								},
							},
							{
								Value: []float64{
									33.4,
								},
							},
							{
								Value: []float64{
									8.3,
									44,
								},
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

func Test_Value_Verify_True(t *testing.T) {
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
								Value: []float64{
									0,
								},
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
								Value: []float64{
									0,
								},
							},
							{
								Value: []float64{
									26,
								},
							},
							{
								Value: []float64{
									99.5,
								},
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
								Value: []float64{
									83,
									83,
									83,
								},
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
								Value: []float64{
									83,
									83,
									83,
								},
							},
							{
								Value: []float64{
									83,
									15,
									8.3,
								},
							},
							{
								Value: []float64{
									2000,
									33.4,
									83,
								},
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
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
