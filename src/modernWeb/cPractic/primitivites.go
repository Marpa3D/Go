// fork-join model
package main

import (
	"fmt"
)

func main() {
	// Создан канал для общения между горутиной(анонимная функция) и главной горутиной main()
	myChannel := make(chan string)
	go func() {
		myChannel <- "data"
	}()
	msg := <-myChannel
	fmt.Println(msg)

	fmt.Println("Финальная функция")
}
