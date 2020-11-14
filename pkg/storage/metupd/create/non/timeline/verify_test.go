package timeline

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	loggerfake "github.com/xh3b4sd/logger/fake"
	"github.com/xh3b4sd/redigo"
	redigofake "github.com/xh3b4sd/redigo/fake"
)

func Test_Timeline_Verify_Bool_False(t *testing.T) {
	testCases := []struct {
		obj *metupd.CreateI
	}{
		// Case 0 ensures that create input without any information provided is
		// not valid.
		{
			obj: &metupd.CreateI{},
		},
		// Case 1 ensures that create input without text is not valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					85,
				},
				Timeline: "tml-al9qy",
			},
		},
		// Case 2 ensures that create input without timeline is not valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					85,
				},
				Text: "Lorem ipsum ...",
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

			ok, err := tml.Verify(tc.obj)
			if err != nil {
				t.Fatal(err)
			}

			if ok != false {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}

func Test_Timeline_Verify_Bool_True(t *testing.T) {
	testCases := []struct {
		obj *metupd.CreateI
	}{
		// Case 0 ensures that create input with only a single datapoint is
		// valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
		},
		// Case 1 ensures that create input with multiple datapoints is valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					85,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
		},
		// Case 2 ensures that create input with multiple datapoints is valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					556,
					1,
					2500,
				},
				Text:     "foo bar #hashtag",
				Timeline: "tml-i45kj",
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

			ok, err := tml.Verify(tc.obj)
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
		obj        *metupd.CreateI
		searchFake func() ([]string, error)
	}{
		// Case 0 ensures that create input with too many y axis coordinates is
		// not valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					85,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
			searchFake: func() ([]string, error) {
				return []string{"1,2"}, nil
			},
		},
		// Case 1 ensures that create input with too many y axis coordinates is
		// not valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					23,
					93,
					53,
					12,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
			searchFake: func() ([]string, error) {
				return []string{"1,2"}, nil
			},
		},
		// Case 2 ensures that create input with too few y axis coordinates is
		// not valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
			searchFake: func() ([]string, error) {
				return []string{"1,2,3,4"}, nil
			},
		},
		// Case 3 ensures that create input with too few y axis coordinates is
		// not valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					85,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
			searchFake: func() ([]string, error) {
				return []string{"1,2,3,4"}, nil
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

			ok, err := tml.Verify(tc.obj)
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
		obj        *metupd.CreateI
		searchFake func() ([]string, error)
	}{
		// Case 0 ensures that create input with the correct amount of y axis
		// coordinates is valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					32,
					85,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
			searchFake: func() ([]string, error) {
				return []string{"1,2,3"}, nil
			},
		},
		// Case 1 ensures that create input with the correct amount of y axis
		// coordinates is valid.
		{
			obj: &metupd.CreateI{
				Yaxis: []int64{
					100,
					150,
				},
				Text:     "Lorem ipsum ...",
				Timeline: "tml-al9qy",
			},
			searchFake: func() ([]string, error) {
				return []string{"1,2,3"}, nil
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

			ok, err := tml.Verify(tc.obj)
			if err != nil {
				t.Fatal(err)
			}

			if ok != true {
				t.Fatalf("\n\n%s\n", cmp.Diff(ok, false))
			}
		})
	}
}
