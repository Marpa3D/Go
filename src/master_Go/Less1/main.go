// Практика использования команды go [args] и горутины
package main

import (
	"fmt"
	"time"
)

func main() {

	// GOOS=windows GOARCH=amd64 go build -o winapp.exe Компиляция под Windows
	// GOOS=linux GOARCH=amd64 go build -o Компиляция под Linux
	// GOOS=darwin GOARCH=amd64 go build -o Компиляция под Mac

	go fmt.Println("Это горутина. Программа работает!)")
	fmt.Println("Обычная инструкция)")

	time.Sleep(2 * time.Second)
}
