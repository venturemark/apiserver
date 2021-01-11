package element

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, []string, error) {
	l := strings.Split(str, ",")

	var t float64
	{
		i, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", nil, tracer.Mask(err)
		}

		t = float64(i)
	}

	var n string
	{
		nam, err := base64.StdEncoding.DecodeString(l[1])
		if err != nil {
			return 0, "", nil, tracer.Mask(err)
		}
		n = string(nam)
	}

	var u []string
	{
		usr, err := base64.StdEncoding.DecodeString(l[2])
		if err != nil {
			return 0, "", nil, tracer.Mask(err)
		}
		u = strings.Split(string(usr), ",")
	}

	return t, n, u, nil
}
