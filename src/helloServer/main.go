// Web сервер
package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Welcome to my site!!! </h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	//time.Sleep(3 * time.Second)
	fmt.Println("Старт сервера. Порт: 8080...")

	// Запуск сервера
	http.ListenAndServe(":8080", nil)
}
