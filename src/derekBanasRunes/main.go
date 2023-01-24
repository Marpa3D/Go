// Курс от Derek Banas. Runes
package main

import (
	"fmt"
	"unicode/utf8"
)

var plf = fmt.Printf

func main() {
	str := "абвгд"
	plf("Счетчик рун: %d\n", utf8.RuneCountInString(str))

	for i, v := range str {
		plf("%d : %#U : %c\n", i, v, v)
	}

}
