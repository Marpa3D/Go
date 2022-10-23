// Практика defer
package main

import "fmt"

func main() {
	for i := 0; i < 8; i++ {
		defer fmt.Println("Отложенный:", -i)
		fmt.Println("Обычный:", i)
	}
}
