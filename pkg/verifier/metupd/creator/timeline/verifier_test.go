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
		fakeScore func() (bool, error)
	}{
		// Case 0 ensures that empty create input is not valid.
		{
			req: &metupd.CreateI{},
			fakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 1 ensures that create input without metadata is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{},
			},
			fakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 2 ensures that create input without timeline ID is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
					},
				},
			},
			fakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 3 ensures that create input without audience ID is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
					},
				},
			},
			fakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 2 ensures that create input for a timeline that does not exist
		// is not valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "0",
					},
				},
			},
			fakeScore: func() (bool, error) {
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
						FakeSorted: func() redigo.Sorted {
							return &fake.Sorted{
								FakeExists: func() redigo.SortedExists {
									return &fake.SortedExists{
										FakeScore: tc.fakeScore,
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
		fakeScore func() (bool, error)
	}{
		// Case 0 ensures that create input is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "1606329189",
					},
				},
			},
			fakeScore: func() (bool, error) {
				return true, nil
			},
		},
		// Case 1 ensures that create input is valid.
		{
			req: &metupd.CreateI{
				Obj: &metupd.CreateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-w4ndz",
						metadata.TimelineID: "1605559909",
					},
				},
			},
			fakeScore: func() (bool, error) {
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
						FakeSorted: func() redigo.Sorted {
							return &fake.Sorted{
								FakeExists: func() redigo.SortedExists {
									return &fake.SortedExists{
										FakeScore: tc.fakeScore,
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
