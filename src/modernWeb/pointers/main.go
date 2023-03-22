package main

import "fmt"

func main() {
	var myString string
	fmt.Printf("Starting data: Adress: %p, Type: %T, Value: %#v\n", &myString, myString, myString)

}
