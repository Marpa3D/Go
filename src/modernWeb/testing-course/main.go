// Тестирование
package main

import (
	"errors"
	"log"
)

func main() {
	res, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Результат деления:", res)
}

func divide(x, y float64) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("Ошибка деления на ноль!")
	}
	result = x / y
	return result, nil
}
