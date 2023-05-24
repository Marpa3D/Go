// Анонимные функции
package main

import "fmt"

func main() {
	// Anonymous function - это функция, котороая не имеет имени
	//и объявляется в строке с помощью литерала функции.
	// Вызывается "прямо сейчас", т.к. у нее нет имени, чтобы вызвать позже!
	func(msg string) {
		fmt.Println(msg)
	}("Это анонимная функция.")
}
