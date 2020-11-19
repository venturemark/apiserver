package data

import (
	"sort"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, []Interface, error) {
	spl := strings.Split(str, ":")

	var uni float64
	{
		i, err := strconv.Atoi(spl[0])
		if err != nil {
			return 0, nil, tracer.Mask(err)
		}

		uni = float64(i)
	}

	var val []Interface
	for _, s := range spl[1:] {
		spl := strings.Split(s, ",")

		{
			w := Wrapper{
				Space: spl[0],
			}

			for _, s := range spl[1:] {
				i, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return 0, nil, tracer.Mask(err)
				}

				w.Value = append(w.Value, i)
			}

			val = append(val, w)
		}
	}

	if len(val) != 0 {
		w := Wrapper{
			Space: "t",
			Value: repeat(uni, len(val[0].GetValue())),
		}

		val = append(val, w)
	}

	sort.Slice(val, func(i, j int) bool {
		return val[i].GetSpace() < val[j].GetSpace()
	})

	return uni, val, nil
}

func repeat(uni float64, num int) []float64 {
	var l []float64

	for i := 0; i < num; i++ {
		l = append(l, uni)
	}

	return l
}
