package main

import (
	"fmt"
	"log"
	"modernWeb/createModule/greetings"
)

func main() {
	msg, err := greetings.Hello("")

	if err != nil {
		log.Println(err)
	}
	fmt.Println(msg)

	msg, err = greetings.Hello("Mark")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(msg)
}
