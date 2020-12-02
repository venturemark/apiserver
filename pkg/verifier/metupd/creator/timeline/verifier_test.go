package timeline

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Timeline_Verify_False(t *testing.T) {
	testCases := []struct {
		req       *metupd.CreateI
		FakeScore func() (bool, error)
	}{
		// Case 0 ensures that empty create input is not valid.
		{
			req: &metupd.CreateI{},
			FakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 1 ensures that create input without metadata is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{},
			},
			FakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 2 ensures that create input without timeline ID is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.UserID: "usr-al9qy",
					},
				},
			},
			FakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 3 ensures that create input without user ID is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
					},
				},
			},
			FakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 2 ensures that create input for a timeline that does not exist
		// is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "0",
						metadata.UserID:     "usr-al9qy",
					},
				},
			},
			FakeScore: func() (bool, error) {
				return false, nil
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
								FakeExists: func() redigo.SortedExists {
									return &fake.SortedExists{
										FakeScore: tc.FakeScore,
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

func Test_Timeline_Verify_True(t *testing.T) {
	testCases := []struct {
		req       *metupd.CreateI
		FakeScore func() (bool, error)
	}{
		// Case 0 ensures that create input is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UserID:     "usr-al9qy",
					},
				},
			},
			FakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 1 ensures that create input is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1605559909",
						metadata.UserID:     "usr-w4ndz",
					},
				},
			},
			FakeScore: func() (bool, error) {
				return true, nil
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
								FakeExists: func() redigo.SortedExists {
									return &fake.SortedExists{
										FakeScore: tc.FakeScore,
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
