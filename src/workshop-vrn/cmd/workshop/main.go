// Web service with Iliya Kaznacheev
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"workshop-vrn/internal/handler"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Get("/hello", h.Hello)
	log.Print("Server starting. Port :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
		log.Print("shutting server down...")
	}

}
