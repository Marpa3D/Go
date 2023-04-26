// Web. Текстовые шаблоны. Синтасический анализ
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("./templates/tpl.gohtml")
	if err != nil {
		log.Println("Ошибка парсинга файла")
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
