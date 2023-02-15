// Добавление и копирование срезов
package main

import "fmt"

func main() {
	// Append - создает новый экземпляр среза, а не изменяет базовый!!!
	// Поэтому нужно снова привязывать к той же переменной
	numbers := []float64{1.0, 4.5}
	fmt.Printf("Adress slice: %p\n", &numbers)
	numbers = append(numbers, 18.0)
	fmt.Println(numbers)
	fmt.Printf("Adress slice: %p\n", &numbers) // Адресс остался тот же

	// Добавление елементов одного среза к другому
	slice := []float64{10.5, 22.8}
	numbers = append(numbers, slice...)
	fmt.Println(numbers)
}
