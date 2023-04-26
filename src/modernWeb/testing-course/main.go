// Тестирование
package main

import "fmt"

const englishHelloPrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "Mark"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
