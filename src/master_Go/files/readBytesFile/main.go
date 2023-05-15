// Считывание n байт из файла Считывание файла с помощью буферизованного считывателя
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Как читать файлы из программы Go
func main() {
	// 1. Октрыть файл для чтения и установить Операционную Систему
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 2. Для примера - считаем из файла только байтовый фрагмент (12 bytes)
	slice := make([]byte, 12)

	numberBytesRead, err := io.ReadFull(file, slice)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Прочитано байт: %d\n", numberBytesRead)
	log.Printf("Прочитанные данные: %s", slice)

	// 3. Читать ве содержимое файла в байтовый слайс (срез)
	file, err = os.Open("main.go")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Содержимое файла: %s", data)

	fmt.Println(strings.Repeat("*", 22))
	fmt.Println("Конец программы")
}
