package element

import (
	"encoding/base64"
	"fmt"
)

func Join(uni float64, des string, nam string, sta string) string {
	return fmt.Sprintf(
		"%d,%s,%s,%s",
		int64(uni),
		string(base64.StdEncoding.EncodeToString([]byte(des))),
		string(base64.StdEncoding.EncodeToString([]byte(nam))),
		string(base64.StdEncoding.EncodeToString([]byte(sta))),
	)
}
