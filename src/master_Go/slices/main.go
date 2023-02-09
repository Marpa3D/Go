// Срезы (слайсы)
package main

import "fmt"

func main() {
	// Обявления среза
	var slice []int // zero value: nil!!! (nil - еще не инициализировано значением)
	fmt.Printf("Type slice: %T, zero value: %#v, len: %+v\n", slice, slice, len(slice))
	//slice[1] = 18
	//fmt.Println(slice) // panic : index out of range [1] length 0

}
