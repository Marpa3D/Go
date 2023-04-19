// fork-join model
package main

import (
	"fmt"
)

func main() {
	// Создан канал для общения между горутиной(анонимная функция) и главной горутиной main()
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "anotherData"
	}()
	// select
	select {
	case msgMyChannel := <-myChannel:
		fmt.Println(msgMyChannel)
	case msgAnotherChannel := <-anotherChannel:
		fmt.Println(msgAnotherChannel)
	}
	fmt.Println("Финальная функция")
}
