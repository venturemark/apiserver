package data

import "fmt"

func Join(uni float64, tex string) string {
	return fmt.Sprintf("%f,%s", uni, tex)
}
