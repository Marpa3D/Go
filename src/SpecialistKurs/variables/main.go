// переменные
package main

import "fmt"

func main() {
	// Два глобальных способа обявления переменных
	// 1.  var variables_name type = value (или значение по умолчанию)
	// 2. Короткое обьявление внутри функции (локальная переменная)
	//    variables_name := value
	var (
		val1 = 188
		val2 = "Привет!)"
		val3 = 124.8
	)
	fmt.Printf("Значение перменной 1: %d\n"+"Тип переменной 1: %T\n", val1, val1)
	fmt.Printf("Значение перменной 2: %s\n"+"Тип переменной 2: %T\n", val2, val2)
	fmt.Printf("Значение перменной 3: %f\n"+"Тип переменной 3: %T", val3, val3)

}
