package main

import (
	"fmt"
	"net/http"
)

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
