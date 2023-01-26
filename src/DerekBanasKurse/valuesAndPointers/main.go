// Семантика значений и указателей
package main

import "fmt"

func main() {
	//Объявляем переменную типа int, привязываем значение 10 к ней
	count := 10
	fmt.Println("count:\tЗначение до передачи в функцию:[", count, "] Адрес в памяти:[", &count, "]")

	// Семантика значений. Передаем копию, а не оригинал пременной в функцию!
	increment(count)
	fmt.Println("count:\tЗначение count после вызова функции:[", count, "] Адрес в памяти:[", &count, "]")

}
func increment(c int) {
	c++
	fmt.Println("conunt:\tЗначение внутри функции:[", c, "] Адрес:[", &c, "]")
}
