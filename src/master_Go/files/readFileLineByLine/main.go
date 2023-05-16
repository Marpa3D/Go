// Чтение файла построчно с помощью сканера
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создаем сканнер
	scanner := bufio.NewScanner(file)
	sucess := scanner.Scan()

	if sucess == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Сканирование завершено и сканнер дошел до конца.")
		} else {
			log.Fatal(err)
		}
	}
	// Чтение первой строки в файле
	fmt.Println("Найдена первая строка:", scanner.Text())

	// Чтение всех строк в файле
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
