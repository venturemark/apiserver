package element

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, string, error) {
	l := strings.Split(str, ",")

	var t float64
	if len(l) >= 1 {
		i, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", "", tracer.Mask(err)
		}

		t = float64(i)
	}

	var m string
	if len(l) >= 2 {
		mes, err := base64.StdEncoding.DecodeString(l[1])
		if err != nil {
			return 0, "", "", tracer.Mask(err)
		}
		m = string(mes)
	}

	var r string
	if len(l) >= 3 {
		rid, err := base64.StdEncoding.DecodeString(l[2])
		if err != nil {
			return 0, "", "", tracer.Mask(err)
		}
		r = string(rid)
	}

	return t, m, r, nil
}
