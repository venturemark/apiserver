package data

import (
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (int64, []Interface, error) {
	spl := strings.Split(str, ":")

	var uni int64
	{
		i, err := strconv.Atoi(spl[0])
		if err != nil {
			return 0, nil, tracer.Mask(err)
		}

		uni = int64(i)
	}

	var val []Interface
	for _, s := range spl[1:] {
		spl := strings.Split(s, ",")

		{
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

			val = append(val, w)
		}

		{
			w := wrapper{
				space: "t",
				value: repeat(uni, len(spl[1:])),
			}

			val = append(val, w)
		}
	}

	return uni, val, nil
}

func repeat(uni int64, num int) []int64 {
	var l []int64

	for i := 0; i < num; i++ {
		l = append(l, uni)
	}

	return l
}
