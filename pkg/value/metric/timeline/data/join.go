package data

import (
	"strconv"
	"strings"
)

func Join(now int64, dat []Interface) string {
	l := []string{
		strconv.Itoa(int(now)),
	}

	for _, d := range dat {
		s := []string{
			d.GetSpace(),
		}

		for _, d := range d.GetValue() {
			l = append(l, strconv.Itoa(int(d)))
		}

		l = append(l, strings.Join(s, ","))
	}

	return strings.Join(l, ":")
}
