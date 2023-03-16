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

	file, err = os.OpenFile("info.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal()
	}
	file.Close()
}
