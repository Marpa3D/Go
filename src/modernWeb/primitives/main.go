// Шаблоны параллелизма
package main

import (
	"fmt"
	"log"
)

func someFunc(num string) {
	fmt.Println("Поток функции someFunc() - ", num)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	go someFunc("4")
	fmt.Println("Основной поток main()")

	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()
	msg := <-myChannel
	log.Println(msg)
}
