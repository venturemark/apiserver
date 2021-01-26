package element

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, string, string, error) {
	l := strings.Split(str, ",")

	var uid float64
	if len(l) >= 1 {
		a, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", "", "", tracer.Mask(err)
		}

		uid = float64(a)
	}

	var oid string
	if len(l) >= 2 {
		n, err := base64.StdEncoding.DecodeString(l[1])
		if err != nil {
			return 0, "", "", "", tracer.Mask(err)
		}
		oid = string(n)
	}

	var tex string
	if len(l) >= 3 {
		t, err := base64.StdEncoding.DecodeString(l[2])
		if err != nil {
			return 0, "", "", "", tracer.Mask(err)
		}
		tex = string(t)
	}

	var usr string
	if len(l) >= 4 {
		u, err := base64.StdEncoding.DecodeString(l[3])
		if err != nil {
			return 0, "", "", "", tracer.Mask(err)
		}
		usr = string(u)
	}

	return uid, oid, tex, usr, nil
}
