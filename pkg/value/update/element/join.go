package element

import "fmt"

func Join(uni float64, tex string) string {
	return fmt.Sprintf("%d,%s", int64(uni), tex)
}
