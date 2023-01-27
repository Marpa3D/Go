// Реализация команды echo в Unix
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, space string
	for i := 1; i < len(os.Args); i++ {
		s += space + os.Args[i]
		space = " "
	}
	if len(os.Args) == 1 {
		fmt.Println("Вы не добавлили аргументы в строке!)\nПерезапустите программу.")
	} else {
		fmt.Println("Аргументы командной строки:", s)
	}

}
