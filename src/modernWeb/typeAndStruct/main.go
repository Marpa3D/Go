// Типы и структуры
package main

import (
	"log"
	"time"
	"unsafe"
)

type Person struct {
	Name      string
	LastName  string
	Phone     string
	Age       int
	BirthDate time.Time
}

var S string

func main() {

	log.Println("Size of type string: ", unsafe.Sizeof(S), "bytes")
	log.Printf("Struct Person.\nAdress in memory: %p, Fields and values:%#v\n", &Person{}, Person{})
	p := Person{
		Name:     "Jason",
		LastName: "Morgan",
		Phone:    "+7968787878",
		Age:      44,
	}
	log.Printf("Struct Person.\nAdress in memory: %p, Fields and values:%#v\n", &p, p)
	log.Println("Size of struct Person: ", unsafe.Sizeof(Person{}), "bytes")

}
