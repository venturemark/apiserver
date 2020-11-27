package element

import "fmt"

func Join(uni float64, nam string) string {
	return fmt.Sprintf("%d,%s", int64(uni), nam)
}
