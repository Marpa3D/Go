// Привет мир от сервера)
package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

// add - сложение двух целых чисел и возвращение их суммы
func add(x, y int) int {
	return x + y
}

// Home функция обработчик страницы "Главная"
func Home(w http.ResponseWriter, r *http.Request) {
	sum := add(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("<h1>Это домашняя страница, а 2 + 2 = %d</h1>", sum))
}

// About - функция обработчик страницы "О нас"
func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Это страница \"о нас\"</h1>")
}

// main - точка входа в приложение
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Сервер запускается...")
	_ = http.ListenAndServe(portNumber, nil)

}
