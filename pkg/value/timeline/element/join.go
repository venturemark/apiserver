package element

import "fmt"

func Join(uni float64, nam string) string {
	return fmt.Sprintf("%f,%s", uni, nam)
}
