package main

import (
	"fmt"
	"modernWeb/createModule/greetings"
)

func main() {
	msg := greetings.Hello("Mark")
	fmt.Println(msg)
}
