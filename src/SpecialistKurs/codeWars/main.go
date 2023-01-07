// Пакет strings - вставка пробелов между символами строки
package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Hello WORLD!!!"
	v := strings.SplitAfter(s, "")
	fmt.Println(v)
}
