// Web. Текстовые шаблоны. Синтасический анализ
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("./templates/tpl.gohtml")
	if err != nil {
		log.Println("Ошибка парсинга файла")
	}
	fmt.Printf("Type tpl: %T\n, Value: %+v", tpl, tpl)

	newFile, err := os.Create("./templates/index.html")
	if err != nil {
		log.Println("Ошибка создания файла")
	}
	defer newFile.Close()

	err = tpl.Execute(newFile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
