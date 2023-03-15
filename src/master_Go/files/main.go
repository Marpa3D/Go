// Работа с файлами
package main

import (
	"fmt"
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

	// Открытие, закрытие файла
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal()
	}
	file.Close()

	// Информация о файле
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("test.txt")
	p := fmt.Println
	p("File name: ", fileInfo.Name())        // Имя
	p("Size of bytes: ", fileInfo.Size())    // Размер файла
	p("Это директория?: ", fileInfo.IsDir()) // Директория
	p("Разрешения: ", fileInfo.Mode())       // Разрешения

	// проверка на существование файла
	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не существует!")
		}
	}

	// Переименование файла
	oldPath := "test.txt"
	newPath := "newText.doc"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	// Удаление файла
	err = os.Remove("newText.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("newText.doc")
	if err != nil {
		log.Fatal(err)
	}

}
