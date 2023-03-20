// Код, который сериализовывает структуру в json-строку
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json: "Имя"`
	Email string `json: "Почта"`
	Age   int    `json: "-"`
}

type Response struct {
}

func ReadResponse(rawResp string) (Response, error) {

}

func main() {
	// Создаем экземпляр(объект) структуры Person
	man := Person{
		Name:  "Илья",
		Email: "seo@yandex.ru",
		Age:   44,
	}

	jsMan, err := json.Marshal(man) // сериализация структуры в JSON
	if err != nil {
		log.Fatalln("Невозможно сериализовать в JSON")
	}
	fmt.Printf("Объект Man: %v", string(jsMan))
}
