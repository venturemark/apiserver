package element

import (
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, error) {
	l := strings.Split(str, ",")

	var u float64
	{
		i, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", tracer.Mask(err)
		}

		u = float64(i)
	}

	var n string
	{
		n = l[1]
	}

	return u, n, nil
}
