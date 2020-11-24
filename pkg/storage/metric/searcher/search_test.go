package searcher

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metric"
	loggerfake "github.com/xh3b4sd/logger/fake"
	"github.com/xh3b4sd/redigo"
	redigofake "github.com/xh3b4sd/redigo/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Searcher_Search_Redis(t *testing.T) {
	testCases := []struct {
		req *metric.SearchI
		str []string
		res *metric.SearchO
	}{
		// Case 0 ensures the storage response is properly constructed based on
		// the returned elements of the sorted set.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
				},
			},
			str: []string{
				"1605559909:y,23",
				"1605559911:y,85",
			},
			res: &metric.SearchO{
				Obj: []*metric.SearchO_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
							metadata.Unixtime: "1605559909",
						},
						Property: &metric.SearchO_Obj_Property{
							Data: []*metric.SearchO_Obj_Property_Data{
								{
									Space: "t",
									Value: []float64{
										1605559909,
									},
								},
								{
									Space: "y",
									Value: []float64{
										23,
									},
								},
							},
						},
					},
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
							metadata.Unixtime: "1605559911",
						},
						Property: &metric.SearchO_Obj_Property{
							Data: []*metric.SearchO_Obj_Property_Data{
								{
									Space: "t",
									Value: []float64{
										1605559911,
									},
								},
								{
									Space: "y",
									Value: []float64{
										85,
									},
								},
							},
						},
					},
				},
			},
		},
		// Case 1 ensures the storage response is properly constructed based on
		// the returned elements of the sorted set.
		{
			req: &metric.SearchI{
				Obj: []*metric.SearchI_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
						},
					},
				},
			},
			str: []string{
				"1605559909:x,0.5,0.7,600:y,23,8.553,300:z,15.3,1,9040",
			},
			res: &metric.SearchO{
				Obj: []*metric.SearchO_Obj{
					{
						Metadata: map[string]string{
							metadata.Timeline: "tml-al9qy",
							metadata.Unixtime: "1605559909",
						},
						Property: &metric.SearchO_Obj_Property{
							Data: []*metric.SearchO_Obj_Property_Data{
								{
									Space: "t",
									Value: []float64{
										1605559909,
										1605559909,
										1605559909,
									},
								},
								{
									Space: "x",
									Value: []float64{
										0.5,
										0.7,
										600,
									},
								},
								{
									Space: "y",
									Value: []float64{
										23,
										8.553,
										300,
									},
								},
								{
									Space: "z",
									Value: []float64{
										15.3,
										1,
										9040,
									},
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

			var s *Searcher
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: &redigofake.Client{
						ScoredFake: func() redigo.Scored {
							return &redigofake.Scored{
								SearchFake: func() ([]string, error) {
									return tc.str, nil
								},
							}
						},
					},
				}

				s, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			res, err := s.Search(tc.req)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.res, res) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.res.String(), res.String()))
			}
		})
	}
}
