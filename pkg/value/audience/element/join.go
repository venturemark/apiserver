package element

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Join(uni float64, tex string, usr []string) string {
	return fmt.Sprintf(
		"%d,%s,%s",
		int64(uni),
		string(base64.StdEncoding.EncodeToString([]byte(tex))),
		string(base64.StdEncoding.EncodeToString([]byte(strings.Join(usr, ",")))),
	)
}
