package data

import (
	"strconv"
	"strings"
)

func Join(now int64, val []Interface) string {
	l := []string{
		strconv.Itoa(int(now)),
	}

	for _, v := range val {
		s := []string{
			v.GetSpace(),
		}

		for _, v := range v.GetValue() {
			s = append(s, strconv.Itoa(int(v)))
		}

		l = append(l, strings.Join(s, ","))
	}

	return strings.Join(l, ":")
}
