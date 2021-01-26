package element

import (
	"encoding/base64"
	"fmt"
)

func Join(mid float64, oid string, tex string, rid string, usr string) string {
	return fmt.Sprintf(
		"%d,%s,%s,%s,%s",
		int64(mid),
		string(base64.StdEncoding.EncodeToString([]byte(oid))),
		string(base64.StdEncoding.EncodeToString([]byte(tex))),
		string(base64.StdEncoding.EncodeToString([]byte(rid))),
		string(base64.StdEncoding.EncodeToString([]byte(usr))),
	)
}
