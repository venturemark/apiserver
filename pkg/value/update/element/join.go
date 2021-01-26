package element

import (
	"encoding/base64"
	"fmt"
)

func Join(uid float64, oid string, tex string, usr string) string {
	return fmt.Sprintf("%d,%s,%s,%s",
		int64(uid),
		string(base64.StdEncoding.EncodeToString([]byte(oid))),
		string(base64.StdEncoding.EncodeToString([]byte(tex))),
		string(base64.StdEncoding.EncodeToString([]byte(usr))),
	)
}
