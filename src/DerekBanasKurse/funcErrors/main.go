// Обработка ошибок в функции
package main

import "fmt"

func getDiv(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("Вы не можете делить на ноль!")
	} else {
		return x / y, nil
	}
}

func main() {
	fmt.Println(getDiv(14, 0))

}
