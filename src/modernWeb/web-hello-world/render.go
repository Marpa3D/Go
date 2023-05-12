package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//rengerTemplate - функция рендера шаблона по имени
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTmpl, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Ошибка парсинга шаблона", err)
		//return
	}
}
