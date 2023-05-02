// Запись байт в файл. os Write and ioutil write File
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"a.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Запись слайса байт(он изменяем, в отличии от строки)
	byteSlice := []byte("Каждый может тебя чему то научить!")
	byteWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Колличество записанных байт: %d\n ", byteWritten)
}
