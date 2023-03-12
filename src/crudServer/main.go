package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404, не найдена страница", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Метод пока не поддерживается", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Привет!)")
}

func formHendler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Ошибка в ParseForm():%v", err)
		return
	}
	fmt.Fprintf(w, "POST метод успешен!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Имя: %s\nАдрес:%s\n", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHendler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Старт сервера!..\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
