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

	// Многомерные массивы ()
	balances := [2][3]int{ // Matrix: 2 строки на 3 столбца
		{1, 3, 8},
		{4, 8, 10},
	}
	fmt.Println(balances)

	// Массивы с ключами элементов
	arr2 := [4]int{
		1: 10,
		0: 8,
		2: 9,
		3: 18,
	}
	fmt.Println(arr2)

	arrStr := [...]string{
		0:      "Jhon",
		8:      "Mark",
		"July", // ??? Index: 8 + 1 = 9!!!
		22:     "Ornella",
	}
	fmt.Printf("%#v\n", arrStr)
}
