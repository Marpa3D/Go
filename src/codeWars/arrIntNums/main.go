// Завершите функцию,
//которая принимает два целых числа ( a, b, где a < b)
//и верните массив всех целых чисел между входными параметрами, включая их.
package main

import "fmt"

func main() {
	fmt.Printf("%v", Between(-1, 8))
}
func Between(a, b int) []int {

	arr := []int{} // Срез через литерал среза

	for i := a; i <= b; i++ {
		arr = append(arr, i) // наполняем срез входящими значениями
	}
	return arr
}
