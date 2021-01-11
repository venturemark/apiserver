package empty

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/venturemark/apigengo/pkg/pbf/metupd"

	"github.com/venturemark/apiserver/pkg/metadata"
)

func Test_Empty_Verify_False(t *testing.T) {
	testCases := []struct {
		req *metupd.UpdateI
	}{
		// Case 0 ensures that empty update input is not valid.
		{
			req: &metupd.UpdateI{},
		},
		// Case 1 ensures that empty update input is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{},
			},
		},
		// Case 2 ensures that update input without timeline ID is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.MetricID:   "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{},
						},
					},
				},
			},
		},
		// Case 3 ensures that update input without audience ID is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.TimelineID: "1606329189",
						metadata.UpdateID:   "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 4 ensures that empty update input is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{},
				},
			},
		},
		// Case 5 ensures that update input with empty data is not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.MetricID:   "1606329189",
						metadata.TimelineID: "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{},
					},
				},
			},
		},
		// Case 6 ensures that update input with empty data and empty text is
		// not valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.MetricID:   "1606329189",
						metadata.TimelineID: "1606329189",
						metadata.UpdateID:   "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{},
						Text: "",
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var v *Verifier
			{
				c := VerifierConfig{}

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

func Test_Empty_Verify_True(t *testing.T) {
	testCases := []struct {
		req *metupd.UpdateI
	}{
		// Case 0 ensures that update input with data is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.MetricID:   "1606329189",
						metadata.TimelineID: "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{},
						},
					},
				},
			},
		},
		// Case 1 ensures that update input without data is valid, because an
		// update request might only be meant to update the text of a metric
		// update.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.TimelineID: "1606329189",
						metadata.UpdateID:   "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Text: "Lorem ipsum ...",
					},
				},
			},
		},
		// Case 2 ensures that update input with data and text is valid.
		{
			req: &metupd.UpdateI{
				Obj: &metupd.UpdateI_Obj{
					Metadata: map[string]string{
						metadata.AudienceID: "aud-al9qy",
						metadata.MetricID:   "1606329189",
						metadata.TimelineID: "1606329189",
						metadata.UpdateID:   "1606329189",
					},
					Property: &metupd.UpdateI_Obj_Property{
						Data: []*metupd.UpdateI_Obj_Property_Data{
							{},
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

			var v *Verifier
			{
				c := VerifierConfig{}

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
