// 1.Домашняя работа. Создать файл
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
