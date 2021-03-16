package subjectclaim

import (
	"context"
	"testing"
)

func Test_Context_Subject_Claim(t *testing.T) {
	{
		s, ok := FromContext(context.Background())
		if ok {
			t.Fatal("ok must be false")
		}
		if s != "" {
			t.Fatal("user must be empty")
		}
	}

	{
		s, ok := FromContext(NewContext(context.Background(), "test"))
		if !ok {
			t.Fatal("ok must be true")
		}
		if s != "test" {
			t.Fatal("user must be test")
		}
	}
}
