// Практика использования команды go [args] и горутины
package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("Это горутина. Программа работает!)")
	fmt.Println("Обычная инструкция)")

	time.Sleep(2 * time.Second)
}
