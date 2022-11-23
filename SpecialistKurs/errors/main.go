// Создание собственного типа "ошибки"
package main

import (
	"errors"
	"fmt"
)

func returnError(num1, num2 int) error {
	if num1 == num2 {
		err := errors.New("Созданная ошибка в функции returnError()!")
		return err
	} else {
		return nil
	}

}

func main() {

	err := returnError(4, 8)
	if err == nil {
		fmt.Println("returnError() завершилась нормально!")
	} else {
		fmt.Println(err)
	}

	err = returnError(8, 8)
	if err == nil {
		fmt.Println("returnError() завершилась нормально!")
	} else {
		fmt.Println(err)
	}
}
