// Привет мир от сервера)
package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

// Home функция обработчик страницы "Главная"
func Home(w http.ResponseWriter, r *http.Request) {

}

// About - функция обработчик страницы "О нас"
func About(w http.ResponseWriter, r *http.Request) {

}

// main - точка входа в приложение
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Сервер запускается...")
	_ = http.ListenAndServe(portNumber, nil)

}
