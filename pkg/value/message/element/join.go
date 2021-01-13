package element

import (
	"encoding/base64"
	"fmt"
)

func Join(uni float64, tex string, rid string) string {
	return fmt.Sprintf(
		"%d,%s,%s",
		int64(uni),
		string(base64.StdEncoding.EncodeToString([]byte(tex))),
		string(base64.StdEncoding.EncodeToString([]byte(rid))),
	)
}
