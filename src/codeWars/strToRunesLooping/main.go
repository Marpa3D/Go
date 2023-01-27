// Выводится строка,в ней остаются только буквы, стоящие на чётных позициях.
package main

import (
	"fmt"
)

func main() {
	str := "Ток: "
	fmt.Println(Cold(str))
}
func Cold(s string) string {
	rs := []rune(s)

	for i, _ := range rs {
		if i%2 == 0 {
			rs = append(rs, rs[i])
		}
	}
	return string(rs)
}
