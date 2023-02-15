// Добавление и копирование срезов
package main

import "fmt"

func main() {
	// Append - создает новый экземпляр среза, а не изменяет базовый!!!
	// Поэтому нужно снова привязывать к той же переменной
	numbers := []float64{1.0, 4.5}
	numbers = append(numbers, 18.0)
	fmt.Println(numbers)

}
