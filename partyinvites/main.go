package main

import "fmt"

type Rsvp struct {
	Name, Email, Phone string
	WillAtteend bool
}
var response = make([]*Rsvp, 0, 10)

func main()  {
	fmt.Println("Поехали!")
}
