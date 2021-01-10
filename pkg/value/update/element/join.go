package element

import (
	"encoding/base64"
	"fmt"
)

func Join(uni float64, tex string) string {
	return fmt.Sprintf("%d,%s",
		int64(uni),
		string(base64.StdEncoding.EncodeToString([]byte(tex))),
	)
}
