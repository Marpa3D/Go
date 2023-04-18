// Тестирование
package main

import "errors"

func main() {

}

func divide(x, y float64) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("Ошибка деления на ноль!")
	}
	result = x / y
	return result, nil
}
