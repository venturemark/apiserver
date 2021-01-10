package updater

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/texupd"
	loggerfake "github.com/xh3b4sd/logger/fake"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Updater_Update_Redis(t *testing.T) {
	testCases := []struct {
		req       *texupd.UpdateI
		fakeValue func() (bool, error)
		res       *texupd.UpdateO
	}{
		// Case 0 ensures that update input with text causes updates on text.
		{
			req: &texupd.UpdateI{
				Obj: &texupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UpdateID:   "1606329189",
						metadata.UserID:     "usr-al9qy",
					},
					Property: &texupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
			fakeValue: testReturn(true, true),
			res: &texupd.UpdateO{
				Obj: &texupd.UpdateO_Obj{
					Metadata: map[string]string{
						metadata.UpdateStatus: "updated",
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var u *Updater
			{
				c := Config{
					Logger: loggerfake.New(),
					Redigo: &fake.Client{
						SortedFake: func() redigo.Sorted {
							return &fake.Sorted{
								FakeUpdate: func() redigo.SortedUpdate {
									return &fake.SortedUpdate{
										FakeValue: tc.fakeValue,
									}
								},
							}
						},
					},
				}

				u, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			res, err := u.Update(tc.req)
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
