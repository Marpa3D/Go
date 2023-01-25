// Обработка ошибок в функции
package main

import "fmt"

func getDiv(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("Вы не можете делить на ноль!\n")
	} else {
		return x / y, nil
	}
}

// Вариационная функция
func getVariaticSum(num ...int) int {
	var sum int
	for _, val := range num {
		sum += val
		fmt.Printf("%d\t", sum)
	}
	fmt.Println("\n")
	return sum
}

func main() {
	fmt.Println(getDiv(14, 0))

	fmt.Println("Общая сумма аргументов: ", getVariaticSum(1, 2, 3, 4, 7, 8))

}
