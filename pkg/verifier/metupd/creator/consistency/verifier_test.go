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
		req       *metupd.CreateI
		FakeIndex func() ([]string, error)
	}{
		// Case 0 ensures that empty create input is not valid.
		{
			req: &metupd.CreateI{},
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 1 ensures that empty create input is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{},
			},
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 2 ensures that create input without metadata is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 3 ensures that empty create input is not valid.
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 4 ensures that create input with too many datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 5 ensures that create input with too many datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1,2"}, nil
			},
		},
		// Case 6 ensures that create input with too few datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 7 ensures that create input with too few datapoints is
		// not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
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
						SortedFake: func() redigo.Sorted {
							return &fake.Sorted{
								FakeSearch: func() redigo.SortedSearch {
									return &fake.SortedSearch{
										FakeIndex: tc.FakeIndex,
									}
								},
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
		req       *metupd.CreateI
		FakeIndex func() ([]string, error)
	}{
		// Case 0 ensures that create input with the correct amount of
		// datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1,2"}, nil
			},
		},
		// Case 1 ensures that create input with the correct amount of
		// datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 2 ensures that create input with the correct amount of
		// datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
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
			FakeIndex: func() ([]string, error) {
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
						SortedFake: func() redigo.Sorted {
							return &fake.Sorted{
								FakeSearch: func() redigo.SortedSearch {
									return &fake.SortedSearch{
										FakeIndex: tc.FakeIndex,
									}
								},
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
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}
