package element

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Join(aid float64, nam string, tim []string, usr []string) string {
	return fmt.Sprintf(
		"%d,%s,%s,%s",
		int64(aid),
		string(base64.StdEncoding.EncodeToString([]byte(nam))),
		string(base64.StdEncoding.EncodeToString([]byte(strings.Join(tim, ",")))),
		string(base64.StdEncoding.EncodeToString([]byte(strings.Join(usr, ",")))),
	)
}
