package element

import (
	"sort"
	"strconv"
	"strings"
)

func Join(now float64, val []Interface) string {
	l := []string{
		strconv.Itoa(int(now)),
	}

	var cop []Interface
	{
		cop = append(cop, val...)

		sort.Slice(cop, func(i, j int) bool {
			return cop[i].GetSpace() < cop[j].GetSpace()
		})
	}

	for _, c := range cop {
		s := []string{
			c.GetSpace(),
		}

		for _, v := range c.GetValue() {
			s = append(s, strconv.FormatFloat(v, 'f', -1, 64))
		}

		l = append(l, strings.Join(s, ","))
	}

	return strings.Join(l, ":")
}
