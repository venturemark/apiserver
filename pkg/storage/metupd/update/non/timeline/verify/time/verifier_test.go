package time

import (
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Time_Verify_False(t *testing.T) {
	uni := time.Unix(int64(1605025038), 0)

	testCases := []struct {
		req *metupd.UpdateI
		now time.Time
	}{
		// Case 0 ensures that update input without any information provided is
		// not valid.
		{
			req: &metupd.UpdateI{},
			now: uni.Add(3 * time.Minute),
		},
		// Case 1 ensures that update input without any information provided is
		// not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{},
			},
			searchFake: func() ([]string, error) {
				return []string{"0:y,1"}, nil
			},
		},
		// Case 2 ensures that update input without metadata is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{},
				},
			},
			now: uni.Add(3 * time.Minute),
		},
		// Case 3 ensures that update input with too old of a timestamp is not
		// valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Unixtime: toString(uni),
					},
				},
			},
			now: uni.Add(6 * time.Minute),
		},
		// Case 4 ensures that update input with too old of a timestamp is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Unixtime: toString(uni),
					},
				},
			},
			now: uni.Add(11 * time.Minute),
		},
		// Case 5 ensures that update input with too old of a timestamp is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Unixtime: toString(uni),
					},
				},
			},
			now: uni.Add(24 * time.Hour),
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{
					Now: func() time.Time {
						return tc.now
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

func Test_Time_Verify_True(t *testing.T) {
	uni := time.Unix(int64(1605025038), 0)

	testCases := []struct {
		req *metupd.UpdateI
		now time.Time
	}{
		// Case 0 ensures that update input with a fresh timestamp is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Unixtime: toString(uni),
					},
				},
			},
			now: uni.Add(1 * time.Minute),
		},
		// Case 1 ensures that update input with a fresh timestamp is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Unixtime: toString(uni),
					},
				},
			},
			now: uni.Add(3 * time.Minute),
		},
		// Case 2 ensures that update input with a fresh timestamp is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.Unixtime: toString(uni),
					},
				},
			},
			now: uni.Add(4 * time.Minute),
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{
					Now: func() time.Time {
						return tc.now
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

func toString(now time.Time) string {
	return strconv.Itoa(int(now.UTC().Unix()))
}
