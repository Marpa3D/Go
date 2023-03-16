// Запись фрагмента байт в файл
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"a.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Записываем в файл. Способ 1.
	bytSlice := []byte("Этот текст был записан из программы на Go!")
	bytWritten, err := file.Write(bytSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Записано байт: %d\n", bytWritten)

	// Записываем в файл. Способ 2.

}
