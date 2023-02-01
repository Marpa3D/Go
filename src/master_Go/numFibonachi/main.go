// Алгоритмы расчета числа Фибоначчи. Медленный и быстрый
package main

import (
	"fmt"
)

func isCountF(n int) int {
	if n <= 1 {
		return n
	} else {
		return isCountF(n-1) + isCountF(n-2)
	}
}

func fiboNum(n int) int {

	slice := make([]int, n+1)
	if n <= 1 {
		return n
	} else { // f = f(n - 1) + f(n - 2)

		slice[0] = 0
		slice[1] = 1
		for i := 2; i <= n; i++ {
			slice[i] = slice[i-1] + slice[i-2]
		}
	}

	return slice[n]
}

func main() {

	fmt.Println("Число Фибоначчи = 32 через рекурсию равно:", isCountF(32))  // медленный способ
	fmt.Println("Число Фибоначчи = 22 быстрым способом равно:", fiboNum(22)) // быстрый способ

}
