// Срезы (слайсы)
package main

import "fmt"

func main() {
	// Обявления среза
	var slice []int // zero value: nil!!! (nil - еще не инициализировано значением)
	fmt.Printf("Type slice: %T, zero value: %#v, len: %+v\n", slice, slice, len(slice))
	//slice[1] = 18
	//fmt.Println(slice) // panic : index out of range [1] length 0

	// Literal slice
	num := []float64{1.1, 2.4, 8.8, 18.0}
	fmt.Println(num)

	// с помощью встроенной функции make
	numbers := make([]int, 10) // срез инициализирован 10-ю zero value for int: 0
	fmt.Printf("%#v\n", numbers)

	// в связке с произвольным type
	type names []string
	firstName := names{
		"Mark",
		"Ornella",
		"Maria",
		"July",
	}
	fmt.Println(firstName)

	// Итерация с range
	for index, value := range firstName {
		fmt.Printf("Index: %+v Value: %+v\n", index, value)
	}

}
