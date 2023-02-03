// Типы данных в Go
package main

import "fmt"

func main() {
	// INT TYPE
	var i8 int8 = 127 // при -129  & 128 => константа -129 & 128 переполняет тип int8. Диапазон -128 до 127!
	fmt.Printf("%T\n", i8)

	var ui16 uint16 = 65535
	fmt.Printf("%T\n", ui16)

	// FLOAT TYPE
	var f1, f2, f3, f4 float64 = 1.1, 1.4, -1.7, 1.8
	fmt.Printf("%T : %T : %T : %T\n", f1, f2, f3, f4)

	// RUNE TYPE
	var r1 rune = 'C'
	fmt.Printf("Rune type:%T\n", r1)
	fmt.Println("Десятичное значение в ASCII", r1)
	fmt.Printf("Шестнадцатиричное значение 'f': %T\n", r1)

	// BOOL TYPE
	var b bool = true
	fmt.Printf("%T:%+v\n", b, b)

	// STRING TYPE
	var s string = "Hello Go!"
	fmt.Printf("%T:%+v\n", s, s)
}
