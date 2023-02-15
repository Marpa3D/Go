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

	// Copy. Копирование элементов слайса из исходного в целевой
	src := []int{1, 3, 5, 7}
	dst := make([]int, len(src))
	copySlice := copy(dst, src)
	fmt.Println(src, dst, copySlice)

	// Если длина целевого среза len == 0 , ничего не скопировать в него.
	// Если его len < len исходного реза - скопируются только элементы, влезающие в len
	src1 := []int{1, 3, 5}
	dst1 := make([]int, 2) // len(dst1) < len(src1)
	copySlice1 := copy(dst1, src1)
	fmt.Println(src1, dst1, copySlice1)
}
