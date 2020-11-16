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

func Test_Timeline_Update_Redis(t *testing.T) {
	testCases := []struct {
		obj        *metupd.UpdateI
		updateFake func() (bool, error)
		met        bool
		upd        bool
	}{
		// Case 0 ensures that update input with yaxis and text updates yaxis
		// and text each.
		{
			obj: &metupd.UpdateI{
				Yaxis: []int64{
					32,
					85,
				},
				Text:      "Lorem ipsum ...",
				Timeline:  "tml-al9qy",
				Timestamp: 1605025038,
			},
			updateFake: testReturn(true, true),
			met:        true,
			upd:        true,
		},
		// Case 1 ensures that update input with only yaxis updates yaxis only.
		// Note that eventhough the mocked redis client returns true for the
		// text update, we should receive false.
		{
			obj: &metupd.UpdateI{
				Yaxis: []int64{
					32,
					85,
				},
				Timeline:  "tml-al9qy",
				Timestamp: 1605025038,
			},
			updateFake: testReturn(true, true),
			met:        true,
			upd:        false,
		},
		// Case 2 ensures that update input with only text updates text only.
		// Note that eventhough the mocked redis client returns true for the
		// axis update, we should receive false.
		{
			obj: &metupd.UpdateI{
				Text:      "Lorem ipsum ...",
				Timeline:  "tml-al9qy",
				Timestamp: 1605025038,
			},
			updateFake: testReturn(true, true),
			met:        false,
			upd:        true,
		},
		// Case 3 ensures that update input with neither yaxis nor text updates
		// none of these resources. In fact this situation should never happen
		// since it is supposed to be covered by Timeline.Verify. Note that
		// eventhough the mocked redis client returns true for either of the
		// axis and the text update, we should receive false for both cases.
		{
			obj: &metupd.UpdateI{
				Timeline:  "tml-al9qy",
				Timestamp: 1605025038,
			},
			updateFake: testReturn(true, true),
			met:        false,
			upd:        false,
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

			obj, err := tml.Update(tc.obj)
			if err != nil {
				t.Fatal(err)
			}
			if tc.met != obj.Metric {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.met, obj.Metric))
			}
			if tc.upd != obj.Update {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.upd, obj.Metric))
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
