package element

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, string, string, string, error) {
	l := strings.Split(str, ",")

	var mid float64
	if len(l) >= 1 {
		i, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", "", "", "", tracer.Mask(err)
		}

		mid = float64(i)
	}

	var oid string
	if len(l) >= 2 {
		b, err := base64.StdEncoding.DecodeString(l[1])
		if err != nil {
			return 0, "", "", "", "", tracer.Mask(err)
		}
		oid = string(b)
	}

	var tex string
	if len(l) >= 3 {
		b, err := base64.StdEncoding.DecodeString(l[2])
		if err != nil {
			return 0, "", "", "", "", tracer.Mask(err)
		}
		tex = string(b)
	}

	var rid string
	if len(l) >= 4 {
		b, err := base64.StdEncoding.DecodeString(l[3])
		if err != nil {
			return 0, "", "", "", "", tracer.Mask(err)
		}
		rid = string(b)
	}

	var usr string
	if len(l) >= 5 {
		b, err := base64.StdEncoding.DecodeString(l[4])
		if err != nil {
			return 0, "", "", "", "", tracer.Mask(err)
		}
		usr = string(b)
	}

	return mid, oid, tex, rid, usr, nil
}
