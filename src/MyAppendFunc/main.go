// реализация собственной функции  для типа int, append (list []int, element int) []int {...}
package main

import "fmt"

func main() {
	list := make([]int, 4, 4)
	list = Append(list, 1)
	fmt.Println(list, len(list), cap(list))
}
func Append(list []int, element int) []int {
	var result []int
	resultLen := len(list) + 1

	if resultLen <= cap(list) {
		result = list[:resultLen]
	} else {
		resultCap := resultLen
		if resultCap < 2*len(list) {
			resultCap = 2 * len(list)
		}
		result = make([]int, resultLen, resultCap)
		copy(result, list)
	}
	result[len(list)] = element
	return result
}
