// Числа Фибоначчи
package main

import "fmt"

// fibOne - самая простая реализация расчета числа Фибоначчи
func fibOne(position int) int {
	seq := []int{0, 1}
	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}
	fmt.Println(seq)
	return seq[len(seq)-1]
}

// Алгоритм 2. Рекурсия - очень медленно, из-за повторов(дублирования) вызова функции
func fibTwo(n int) int {
	fmt.Println(n)
	if n < 2 {
		return n
	}
	a, b := fibTwo(n-1), fibTwo(n-2)
	return a + b
}

func main() {
	fmt.Println("Способ 1. Ваше число Фибоначчи:", fibOne(8))

	fmt.Println("Способ 2. Ваше число Фибоначчи:", fibTwo(8))

}
