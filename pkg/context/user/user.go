package user

import (
	"context"
)

type t string

var k t = "user"

func NewContext(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(k).(string)
	return v, ok
}
