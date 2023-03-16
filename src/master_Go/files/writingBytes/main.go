// Запись фрагмента байт в файл
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"info.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Записываем в файл. Способ 1.
	bytSlice := []byte("Обновленный текст был записан из программы на Go!")
	bWritten, err := file.Write(bytSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Записано байт: %d\n", bWritten)

	// Записываем в файл. Способ 2. Библиотека io/ioutil
	bs := []byte("Способ записи 2. Текст, записанный с помощбю метода ioutil.WriteFile()")
	err = ioutil.WriteFile("info2.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Считывание содержимого файла.
	rf, err := ioutil.ReadFile("info2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", rf)
}
