// Общение со срезом
package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}
	slice[0] *= 10
	slice[2] = 300
	fmt.Println("Slice value:", slice, len(slice), cap(slice))
	slice = append(slice, 20) // Добавление нового элемента
	fmt.Println("Slice value:", slice, len(slice), cap(slice))
}
