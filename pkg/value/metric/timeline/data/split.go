package data

import (
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (int64, []Interface, error) {
	spl := strings.Split(str, ":")

	var now int64
	{
		i, err := strconv.Atoi(spl[0])
		if err != nil {
			return 0, nil, tracer.Mask(err)
		}

		now = int64(i)
	}

	var dat []Interface
	for _, s := range spl[1:] {
		spl := strings.Split(s, ",")

		w := wrapper{
			space: spl[0],
		}

		for _, s := range spl[1:] {
			i, err := strconv.Atoi(s)
			if err != nil {
				return 0, nil, tracer.Mask(err)
			}

			w.value = append(w.value, int64(i))
		}

		dat = append(dat, w)
	}

	return now, dat, nil
}
