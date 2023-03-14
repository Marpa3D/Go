package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Привет от сервера на Go!!!)</H1>")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
