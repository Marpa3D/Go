// Алгоритмы. Линейный поиск, O(n)
package main

import (
	"fmt"
)

func linearSearch(arr []int, target int) int {

	for i, _ := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1 // ВВозвращаемое значение в стиле псевдокода!)))
}

func main() {
	a := []int{1, 4, 5, 7, 9, 8, 10}
	fmt.Println(linearSearch(a, 8))
}
