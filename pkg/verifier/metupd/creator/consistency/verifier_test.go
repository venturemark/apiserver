package consistency

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Consistency_Verify_False(t *testing.T) {
	testCases := []struct {
		req        *metupd.CreateI
		searchFake func() ([]string, error)
	}{
		// Case 0 ensures that update input with too many datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
									85,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 1 ensures that update input with too many datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
									85,
									1,
									2500,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2"}, nil
			},
		},
		// Case 2 ensures that update input with too few datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 3 ensures that update input with too few datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
									85,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{
					Redigo: &fake.Client{
						ScoredFake: func() redigo.Scored {
							return &fake.Scored{
								SearchFake: tc.searchFake,
							}
						},
					},
				}

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

func Test_Consistency_Verify_True(t *testing.T) {
	testCases := []struct {
		req        *metupd.CreateI
		searchFake func() ([]string, error)
	}{
		// Case 0 ensures that empty update input is valid.
		{
			req: &metupd.CreateI{},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 1 ensures that empty update input is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 2 ensures that update input without data is valid, because an
		// update request might only be meant to update the text of a metric
		// update.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 3 ensures that update input without data is valid, because an
		// update request might only be meant to update the text of a metric
		// update.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 4 ensures that update input with the correct amount of
		// datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
									85,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2"}, nil
			},
		},
		// Case 5 ensures that update input with the correct amount of
		// datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
									85,
									1,
									2500,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 6 ensures that update input with the correct amount of
		// datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Value: []float64{
									32,
									85,
									1,
									2500,
								},
							},
							{
								Value: []float64{
									32,
									85,
									1,
									2500,
								},
							},
						},
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:x,1,2,3,4", "0:y,1,2,3,4"}, nil
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{
					Redigo: &fake.Client{
						ScoredFake: func() redigo.Scored {
							return &fake.Scored{
								SearchFake: tc.searchFake,
							}
						},
					},
				}

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
