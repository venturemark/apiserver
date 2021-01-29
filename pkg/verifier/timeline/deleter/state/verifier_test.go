package state

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/timeline"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pkg/fake"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_State_Verify_False(t *testing.T) {
	testCases := []struct {
		req       *timeline.DeleteI
		fakeScore func() ([]string, error)
	}{
		// Case 0 ensures that empty delete input is not valid.
		{
			req: &timeline.DeleteI{},
			fakeScore: func() ([]string, error) {
				return []string{`{"obj":{"property":{"stat":"archived"}}}`}, nil
			},
		},
		// Case 1 ensures that empty delete input is not valid.
		{
			req: &timeline.DeleteI{
				Obj: &timeline.DeleteI_Obj{},
			},
			fakeScore: func() ([]string, error) {
				return []string{`{"obj":{"property":{"stat":"archived"}}}`}, nil
			},
		},
		// Case 2 ensures that delete input with an active timeline is not
		// valid.
		{
			req: &timeline.DeleteI{
				Obj: &timeline.DeleteI_Obj{
					Metadata: map[string]string{
						metadata.OrganizationID: "<id>",
						metadata.TimelineID:     "1606329189",
					},
				},
			},
			fakeScore: func() ([]string, error) {
				return []string{`{"obj":{"property":{"stat":"active"}}}`}, nil
			},
		},
		// Case 3 ensures that delete input with arbitrary timeline state is not
		// valid.
		{
			req: &timeline.DeleteI{
				Obj: &timeline.DeleteI_Obj{
					Metadata: map[string]string{
						metadata.OrganizationID: "<id>",
						metadata.TimelineID:     "1606329189",
					},
				},
			},
			fakeScore: func() ([]string, error) {
				return []string{`{"obj":{"property":{"stat":"foo"}}}`}, nil
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
								FakeSearch: func() redigo.SortedSearch {
									return &fake.SortedSearch{
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

func Test_State_Verify_True(t *testing.T) {
	testCases := []struct {
		req       *timeline.DeleteI
		fakeScore func() ([]string, error)
	}{
		// Case 0 ensures that delete input with an archived timeline is valid.
		{
			req: &timeline.DeleteI{
				Obj: &timeline.DeleteI_Obj{
					Metadata: map[string]string{
						metadata.OrganizationID: "<id>",
						metadata.TimelineID:     "1606329189",
					},
				},
			},
			fakeScore: func() ([]string, error) {
				return []string{`{"obj":{"property":{"stat":"archived"}}}`}, nil
			},
		},
		// Case 1 ensures that delete input with an archived timeline is valid.
		{
			req: &timeline.DeleteI{
				Obj: &timeline.DeleteI_Obj{
					Metadata: map[string]string{
						metadata.OrganizationID: "<id>",
						metadata.TimelineID:     "1606329189",
					},
				},
			},
			fakeScore: func() ([]string, error) {
				return []string{`{"obj":{"property":{"name":"mrr","stat":"archived"}}}`}, nil
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
								FakeSearch: func() redigo.SortedSearch {
									return &fake.SortedSearch{
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
