// Web сервер
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Привет, это сервер на Go!")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Колличество записанных байт: %d", n))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
