// Чтение и запись JSON
package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Color     string `json:"color"`
	HasCat    bool   `json:"has_cat"`
}

func main() {

	myJson := `
[
	{
		"first_name": "Mark",
		"last_name": "Wolfman",
		"color": "white",
		"has_cat": true
	},
	{
		"first_name": "Jhon",
		"last_name": "Malkovich",
		"color": "red",
		"has_cat": false
	}
]`
	// Преобразование формата JSON в структуру
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Формат не соответствует JSON", err)
	}
	log.Printf("unmarshalled: %v\n", unmarshalled)
}
