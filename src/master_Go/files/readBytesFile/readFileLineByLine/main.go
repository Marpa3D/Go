// Чтение файла построчно с помощью сканера
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
