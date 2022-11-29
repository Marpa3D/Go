// Кодовая точка Юникода (руна) и работа с ней
package main

import "fmt"

func main() {
	slice := []rune{'©', '♬', '♡', '🙂'}

	for i, v := range slice {
		fmt.Printf("\nСимвол (руна): %c, Unicode: %U, Тип: %T, Позиция: %d", v, v, v, i)
	}
}
