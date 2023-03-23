// Методы
package main

import "log"

type myStruct struct {
	FirstName string
}

func (s *myStruct) retFirstName() string {
	return s.FirstName
}

func main() {
	var value myStruct
	value.FirstName = "Mark"
	value1 := myStruct{
		FirstName: "Ilia",
	}

	log.Println("Значение value.FirstName:", value.FirstName)
	log.Println("Значение value1:", value1.FirstName)
	log.Println("Значение value.FirstName через метод:", value.retFirstName())
	log.Println("Значение value.FirstName через метод:", value1.retFirstName())
}
