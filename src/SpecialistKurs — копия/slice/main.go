// Общение со срезом
package main

import "fmt"

func main() {

	// Срез (он же слайс) - это динамическая обвязка над массивом. Это фрагмент. диапазон (область) массива.
	//slice := []int{10, 20, 30, 40}
	//slice[0] *= 10
	//slice[2] = 300
	//fmt.Println("Slice value:", slice, len(slice), cap(slice))
	//slice = append(slice, 20) // Добавление нового элемента
	//fmt.Println("Slice value:", slice, len(slice), cap(slice))
	//
	//// итерирование по срезу
	//for i, v := range slice {
	//	fmt.Println(i, v)
	//}

	startArr := [5]int{10, 20, 30, 40, 50}
	var startSlice []int = startArr[0:3] // Срез инициализируется пустыми квадратными скобками
	fmt.Println(startSlice)
	// 1) Создали срез на уже существующем массиве startArr

	// 2) Создание среза (слайса) без явной инициализации массива
	secondSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("Second slice:", secondSlice)
}
