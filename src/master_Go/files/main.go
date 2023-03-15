// Работа с файлами
package main

import (
	"log"
	"os"
)

func main() {

	var err error
	newFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	newFile.Close()

	// Очистка или обрезка файла
	err = os.Truncate("test.txt", 22) // Обрезать до 22 символов.
	if err != nil {
		log.Fatal(err)
	}

}
