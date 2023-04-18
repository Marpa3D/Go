// Привет мир от сервера)
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := fmt.Fprintf(w, "Привет мир от серевера на Go!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Колличество прочитанных байт: %d", b))
	})
	fmt.Println("Сервер запускается...")
	_ = http.ListenAndServe(":8080", nil)
}
