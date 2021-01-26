package element

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/xh3b4sd/tracer"
)

func Split(str string) (float64, string, []string, []string, error) {
	l := strings.Split(str, ",")

	var aid float64
	if len(l) >= 1 {
		a, err := strconv.ParseFloat(l[0], 64)
		if err != nil {
			return 0, "", nil, nil, tracer.Mask(err)
		}

		aid = float64(a)
	}

	var nam string
	if len(l) >= 2 {
		n, err := base64.StdEncoding.DecodeString(l[1])
		if err != nil {
			return 0, "", nil, nil, tracer.Mask(err)
		}
		nam = string(n)
	}

	var tim []string
	if len(l) >= 3 {
		t, err := base64.StdEncoding.DecodeString(l[2])
		if err != nil {
			return 0, "", nil, nil, tracer.Mask(err)
		}
		if string(t) == "" {
			tim = []string{}
		} else {
			tim = strings.Split(string(t), ",")
		}
	}

	var usr []string
	if len(l) >= 4 {
		u, err := base64.StdEncoding.DecodeString(l[3])
		if err != nil {
			return 0, "", nil, nil, tracer.Mask(err)
		}
		if string(u) == "" {
			usr = []string{}
		} else {
			usr = strings.Split(string(u), ",")
		}
	}

	return aid, nam, tim, usr, nil
}
