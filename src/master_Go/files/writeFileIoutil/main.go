// Запись байт в файл. os Write and ioutil write File
package main

import (
	"io/ioutil"
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

	// 1.Запись слайса байт(он изменяем, в отличии от строки) с помощью os.OpenFile
	byteSlice := []byte("Каждый может тебя чему то научить!")
	byteWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Колличество записанных байт: %d\n ", byteWritten)

	// Создание(если нет) файла и запись в него слайса байт
	bs := []byte("Путь начинается с маленького шага...")
	err = ioutil.WriteFile("b.txt", bs, 0644)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Колличество записанных байт: %d\n ", len(bs))
}
