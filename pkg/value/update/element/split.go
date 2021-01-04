package element

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, error) {
	l := strings.Split(str, ",")

	var n float64
	{
		i, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", tracer.Mask(err)
		}

		n = float64(i)
	}

	var t string
	{
		tex, err := base64.StdEncoding.DecodeString(l[1])
		if err != nil {
			return 0, "", tracer.Mask(err)
		}
		t = string(tex)
	}

	return n, t, nil
}
