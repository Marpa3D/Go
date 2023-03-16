//Переименуйте файл, созданный в упражнении №1, в data.txt
//Проверьте, существует ли файл, прежде чем переименовывать его.
//Если он не существует, выведите сообщение и остановите программу.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// create file
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	fileInfo, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Файл не существует!")
		}
	}
	fmt.Println(fileInfo.Mode())

	// rename file
	old, new := "info.txt", "data.txt"
	err = os.Rename(old, new)
	if err != nil {
		log.Fatal(err)
	}

	//Удалите файл, созданный в упражнении №1.
	//Позаботьтесь о том, чтобы файл теперь вызывался data.txt (он был переименован в упражнении №2).
	//Выполните обработку ошибок.
	err = os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
