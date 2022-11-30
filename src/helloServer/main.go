// Web сервер
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет, это сервер на Go!")
	})
	http.ListenAndServe(":8080", nil)
}
