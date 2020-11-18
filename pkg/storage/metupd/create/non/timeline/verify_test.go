package timeline

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	loggerfake "github.com/xh3b4sd/logger/fake"
	"github.com/xh3b4sd/redigo"
	redigofake "github.com/xh3b4sd/redigo/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Timeline_Verify_Input_False(t *testing.T) {
	testCases := []struct {
		req *metupd.CreateI
	}{
		// Case 0 ensures that create input without any information provided is
		// not valid.
		{
			req: &metupd.CreateI{},
		},
		// Case 1 ensures that create input without metadata is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 2 ensures that create input without timeline is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 3 ensures that create input without text is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
					},
				},
			},
		},
		// Case 4 ensures that create input without dimensional space is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "",
								Value: []int64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 5 ensures that create input without datapoints is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 6 ensures that create input with different amounts of datapoints
		// per dimension is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "x",
								Value: []int64{
									32,
								},
							},
							{
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 7 ensures that create input with a duplicated dimensional space
		// is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
								},
							},
							{
								Space: "y",
								Value: []int64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 8 ensures that create input with the reserved dimensional space
		// t is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "x",
								Value: []int64{
									73,
									91,
								},
							},
							{
								Space: "y",
								Value: []int64{
									22,
									94,
								},
							},
							{
								Space: "t",
								Value: []int64{
									20,
									16,
								},
							},
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

			var tml *Timeline
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: redigofake.New(),
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok, err := tml.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Timeline_Verify_Input_True(t *testing.T) {
	testCases := []struct {
		req *metupd.CreateI
	}{
		// Case 0 ensures that create input with only a single datapoint is
		// valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 1 ensures that create input with multiple datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 2 ensures that create input with multiple datapoints is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
									85,
									1,
									2500,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 3 ensures that create input with multiple dimensional spaces is
		// valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "x",
								Value: []int64{
									32,
									85,
								},
							},
							{
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 4 ensures that create input with multiple dimensional spaces is
		// valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "x",
								Value: []int64{
									73,
									91,
								},
							},
							{
								Space: "y",
								Value: []int64{
									22,
									94,
								},
							},
							{
								Space: "z",
								Value: []int64{
									20,
									16,
								},
							},
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

			var tml *Timeline
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: redigofake.New(),
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok, err := tml.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, true))
			}
		})
	}
}

func Test_Timeline_Verify_Search_False(t *testing.T) {
	testCases := []struct {
		req        *metupd.CreateI
		searchFake func() ([]string, error)
	}{
		// Case 0 ensures that create input with too many y axis coordinates is
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
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 1 ensures that create input with too many y axis coordinates is
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
								Space: "y",
								Value: []int64{
									32,
									85,
									1,
									2500,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2"}, nil
			},
		},
		// Case 2 ensures that create input with too few y axis coordinates is
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
								Space: "y",
								Value: []int64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 3 ensures that create input with too few y axis coordinates is
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
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 4 ensures that create input with dimensional space identified
		// by anything other than a single character is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "toolong",
								Value: []int64{
									32,
								},
							},
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

			var tml *Timeline
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: &redigofake.Client{
						ScoredFake: func() redigo.Scored {
							return &redigofake.Scored{
								SearchFake: tc.searchFake,
							}
						},
					},
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok, err := tml.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Timeline_Verify_Search_True(t *testing.T) {
	testCases := []struct {
		req        *metupd.CreateI
		searchFake func() ([]string, error)
	}{
		// Case 0 ensures that create input with the correct amount of y axis
		// coordinates is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
									85,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2"}, nil
			},
		},
		// Case 1 ensures that create input with the correct amount of y axis
		// coordinates is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []int64{
									32,
									85,
									1,
									2500,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1,2,3,4"}, nil
			},
		},
		// Case 2 ensures that create input with the correct amount of y axis
		// coordinates is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
					},
					Property: &metupd.CreateI_Obj_Property{
						Data: []*metupd.CreateI_Obj_Property_Data{
							{
								Space: "x",
								Value: []int64{
									32,
									85,
									1,
									2500,
								},
							},
							{
								Space: "y",
								Value: []int64{
									32,
									85,
									1,
									2500,
								},
							},
						},
						Text: "Lorem ipsum ...",
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

			var tml *Timeline
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: &redigofake.Client{
						ScoredFake: func() redigo.Scored {
							return &redigofake.Scored{
								SearchFake: tc.searchFake,
							}
						},
					},
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			ok, err := tml.Verify(tc.req)
			if err != nil {
				t.Fatal(err)
			}

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}
