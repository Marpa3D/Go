// Массивы
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Arrays in Go

	s1 := [...]string{"X", "Y", "Z"}
	fmt.Printf("%v\t%#v\t%d\n", s1, s1, len(s1))

	// 1.iteration arr range
	for i, v := range s1 {
		fmt.Printf("Index:%+v\tValue:%+v\n", i, v)
	}

	fmt.Println(strings.Repeat("@", 54)) // создает горизонтальну. строку из символа '@'

	// 2. Iteration arr loop for (c - signature)
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i])
	}
}
