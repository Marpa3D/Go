// Бинарный поиск
package main

import (
	"fmt"
	"sort"
)

// binSearch - ищет нужный элемент в отсортированном срезе, посредством
// последовательного уменьшения объема данных, подлежащих поиску
func binSearch(elem int, data []int) bool {
	sort.Ints(data)

	low := 0
	high := len(data) - 1

	for low < high {
		median := (low + high) / 2 // отбрасывание ненужной половины данных!
		if data[median] < elem {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(data) || data[low] != elem {
		return false
	}
	return true
}

func main() {
	num := 15
	slice := []int{8, 2, 9, 3, 5, 1, 15, 22}
	fmt.Printf("Наличие элемента %d в срезе %#v: %v\n", num, slice, binSearch(num, slice))
}
