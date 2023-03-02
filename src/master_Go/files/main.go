// Работа с файлами
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	pf := fmt.Printf

	var newFile *os.File
	pf("%T\n", newFile)

	var err error
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	newFile.Close()

	// Очистка файла
	err = os.Truncate("test.txt", 0)
}
