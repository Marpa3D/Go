// defer - оператор отсрочки
package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is func foo()")
}

func bar() {
	fmt.Println("This is func bar()")
}

func fooBar() {
	fmt.Println("This is func fooBar()")
}

func main() {
	// defer - откладывает выполнение функции пока текущая функция не завершится
	// он помещает вызов указанной функции в список вызовов,
	//который выполняется после завершения работы окружающей функции
	defer foo()
	bar()
	fmt.Println("Текст после вызовов foo() и bar()")
	defer fooBar() // defer помещает йункции в стек.
	// LIFO (последний пришел, первый ушел)

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}
