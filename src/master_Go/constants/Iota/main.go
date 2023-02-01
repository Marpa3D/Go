// Константы. iota
package main

import "fmt"

func main() {
	// iota - это генератор чисел для констант, начинается с нуля.
	// Это последовательные константы 0, 1, 2, 3,...
	// iota сбрасывается на 0 при встерече const
	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)
	fmt.Println(a1, a2, a3)

	// Краткая запись
	const (
		a11 = iota // 0
		a22        // 1
		a33        // 2
	)
	fmt.Println(a11, a22, a33)

	const (
		North = iota // default 0
		East
		South
		West
	)
	fmt.Println(West) // 3

	// с варажением
	const (
		c1 = (iota * 2) + 1 // 1
		b1                  // 3
		f1                  // (2 * 2) + 1 = 5
	)
	fmt.Println(c1, b1, f1) // 1, 3, 5

	// x = -2, y = -4, z = -5
	const (
		x = -(iota + 2) // -2
		_               // -3
		y               // -4
		z               // -5
	)
	fmt.Println(x, y, z)
}
