// Реализация команды echo в Unix
package main

import (
	"fmt"
	"os"
)

func main() {
	s, space := "", ""
	for i, arg := range os.Args[:] {
		s += space + arg
		space = " "
		fmt.Printf("Индекс строки: [%d]  Аргумент строки: %s\n", i, arg)
	}
	if len(os.Args) == 1 {
		fmt.Println("Вы не добавлили аргументы в строке!)\nПерезапустите программу.")
	} else {
		fmt.Println("Аргументы командной строки:", s)
	}

}
