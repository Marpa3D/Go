package main

import "fmt"

func main() {
	type age int
	var age1 age = 44
	var young age = 44
	var old age = 88
	fmt.Println(age1, young, old)
}
