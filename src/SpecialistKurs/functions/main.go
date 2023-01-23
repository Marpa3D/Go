// Специалист. Функции.
package main

import "fmt"

// 1. Явная функция - именованный блок кода, который может вызываться по необходимости. Имеет имя
// Функцию нужно определить + вызвать.
// Сигнатура функции.
// func nameFunction(param type) returnValues {
//  code...
//   }
// param type и returnValues - опционально
func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Сумма двух целых чисел:", add(4, 4))
}
