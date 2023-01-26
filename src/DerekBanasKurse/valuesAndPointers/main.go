// Семантика значений и указателей
package main

import "fmt"

func incrementValue(c int) {
	c++
	fmt.Println("conunt:\tЗначение внутри функции:[", c, "] Адрес:[", &c, "]")
}

func incrementPointer(c *int) {
	*c++
	fmt.Println("count:\tЗначение по адресу count изменилось:[", *c, "] Адрес в памяти:[", &c, "]")
}

// Два варианта в одной функции
func increment(c int, cp *int) {
	c++
	fmt.Println("conunt:\tЗначение внутри функции:[", c, "] Адрес:[", &c, "]")
	*cp++
	fmt.Println("count:\tЗначение по адресу count изменилось:[", *cp, "] Адрес в памяти:[", &cp, "]")
}

func main() {
	// 1. Семантика значений
	//Объявляем переменную типа int, привязываем значение 10 к ней
	count := 10
	fmt.Println("count:\tЗначение до передачи в функцию:[", count, "] Адрес в памяти:[", &count, "]")

	// Семантика значений. Передаем копию, а не оригинал пременной в функцию!
	incrementValue(count)
	fmt.Println("count:\tЗначение count после вызова функции:[", count, "] Адрес в памяти:[", &count, "]")

	// А здесь, через указатель, меняем само значение переменной count
	incrementPointer(&count)

	// Объединение обоих вариантов в одной функции
	increment(count, &count)

}
