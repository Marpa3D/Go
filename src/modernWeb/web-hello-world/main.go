// Привет мир от сервера)
package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

// main - точка входа в приложение
func main() {
	http.HandleFunc("/", pathHendler)
	//http.HandleFunc("/about", About)

	fmt.Println("Сервер запускается в порту:8080")
	_ = http.ListenAndServe(portNumber, nil)

}
