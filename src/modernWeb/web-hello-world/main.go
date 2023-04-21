// Привет мир от сервера)
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber string = ":8080"

// Home функция обработчик страницы "Главная"
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset = utf-8")
	renderTemplate(w, "home.page.tmpl")
}

// About - функция обработчик страницы "О нас"
func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset = utf-8")
	renderTemplate(w, "about.page.tmpl")
}

//rengerTemplate - функция рендера шаблона по имени
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTmpl, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Ошибка парсинга шаблона", err)
		//return
	}
}

func pathHendler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		Home(w, r)
	case "/about":
		About(w, r)
	default:
		fmt.Fprintf(w, "<h1>404 - такой страницы не существует</h1>")
	}
}

// main - точка входа в приложение
func main() {
	http.HandleFunc("/", pathHendler)
	//http.HandleFunc("/about", About)

	fmt.Println("Сервер запускается в порту:8080")
	_ = http.ListenAndServe(portNumber, nil)

}
