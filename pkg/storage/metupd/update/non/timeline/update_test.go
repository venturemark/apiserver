package timeline

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	loggerfake "github.com/xh3b4sd/logger/fake"
	"github.com/xh3b4sd/redigo"
	redigofake "github.com/xh3b4sd/redigo/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Timeline_Update_Redis(t *testing.T) {
	testCases := []struct {
		req        *metupd.UpdateI
		updateFake func() (bool, error)
		res        *metupd.UpdateO
	}{
		// Case 0 ensures that update input with data and text causes updates on
		// data and text each.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
						metadata.Unixtime: "1605025038",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []float64{
									32,
								},
							},
						},
						Text: "Lorem ipsum ...",
					},
				},
			},
			updateFake: testReturn(true, true),
			res: &metupd.UpdateO{
				Obj: &metupd.UpdateO_Obj{
					Metadata: map[string]string{
						metadata.MetricStatus: "updated",
						metadata.UpdateStatus: "updated",
					},
				},
			},
		},
		// Case 1 ensures that update input with only data causes updates on
		// data only. Note that eventhough the mocked redis client returns true
		// for the text update, we should receive false.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
						metadata.Unixtime: "1605025038",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{
								Space: "y",
								Value: []float64{
									32,
								},
							},
						},
					},
				},
			},
			updateFake: testReturn(true, true),
			res: &metupd.UpdateO{
				Obj: &metupd.UpdateO_Obj{
					Metadata: map[string]string{
						metadata.MetricStatus: "updated",
					},
				},
			},
		},
		// Case 2 ensures that update input with only text causes updates on
		// text only. Note that eventhough the mocked redis client returns true
		// for the axis update, we should receive false.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
						metadata.Unixtime: "1605025038",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
			updateFake: testReturn(true, true),
			res: &metupd.UpdateO{
				Obj: &metupd.UpdateO_Obj{
					Metadata: map[string]string{
						metadata.UpdateStatus: "updated",
					},
				},
			},
		},
		// Case 3 ensures that update input with data nor text does not cause
		// updates of any of these resources. In fact this situation should
		// never happen since it is supposed to be covered by Timeline.Verify.
		// Note that eventhough the mocked redis client returns true for either
		// of the axis and the text update, we should receive false for both
		// cases.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Timeline: "tml-al9qy",
						metadata.Unixtime: "1605025038",
					},
					Property: &metupd.UpdateI_Obj_Property{},
				},
			},
			updateFake: testReturn(true, true),
			res: &metupd.UpdateO{
				Obj: &metupd.UpdateO_Obj{
					Metadata: map[string]string{},
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
								UpdateFake: tc.updateFake,
							}
						},
					},
				}

				tml, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			res, err := tml.Update(tc.req)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.res, res) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.res, res))
			}
		})
	}
}

func testReturn(met bool, upd bool) func() (bool, error) {
	var c int

	return func() (bool, error) {
		if c == 0 {
			c++
			return met, nil
		}

		return upd, nil
	}
}
