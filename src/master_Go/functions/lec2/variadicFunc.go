// Вариативные функции
package main

import (
	"fmt"
	"strings"
)

// Функция - это фрагмент кода для выполнения определенной задачи на основании входящих аргументов
// Вариационные функции - это функции, которые принимают переменное число аргументов
// Префикс или три точки ... перед типом параметра функции делают ее вариационной!
// Если функция имеет праметры разных типов, то только последний тип может быть вариативным.
// Когда применять?

// 1. Когда колличество параметров не известно
// 2. Не хотите создавать временный срез только для того, чтобы передать его функции
// 3. Для удобочитаемости кода
func f1(n ...int) {
	fmt.Printf("Тип данных: %T\n", n)
	fmt.Printf("%#v\n", n)
}

func f2(a ...int) {
	a[1] = 22
}

func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.0
	product := 1.0
	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, " ")

	resString := fmt.Sprintf("Age: %d. Full Name: %s", age, fullName)

	return resString
}

func main() {
	f1()
	f1(1, 2, 3, 4, 5, 6, 7, 8)

	// Пример. Вариативная функция append.
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)

	f1(nums...)
	f2(nums...) // изменяет переменную nums не создавая копию в памяти!!!
	fmt.Printf("Nums: %d\n", nums)

	fmt.Println(sumAndProduct(4.0, 4.0, 8., 19.0))
	a, b := sumAndProduct(3., 7., 8.)
	fmt.Println(a, b)

	fmt.Println(personInfo(44, "Shtolle", "Shwarchold"))
}
