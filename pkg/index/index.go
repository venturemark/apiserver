package index

import (
	"fmt"
	"strings"
)

const (
	Name = "name:%s"
)

func New(f string, v ...interface{}) string {
	var l []interface{}

	for _, e := range v {
		s, ok := e.(string)
		if ok {
			l = append(l, sanitize(s))
		} else {
			l = append(l, e)
		}
	}

	return fmt.Sprintf(f, l...)
}

func sanitize(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}
