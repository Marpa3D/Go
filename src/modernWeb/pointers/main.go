package main

import (
	"fmt"
	"log"
)

// Функция просто меняет значение переменной по указанному адресу ее нахождения
func changeUserPointer(s *string) {
	log.Println("Adress s:", s)
	newValue := "Красный"
	*s = newValue
}

func main() {

	var myString string
	fmt.Printf("Starting data: Adress: %p, Type: %T, Value: %#v\n", &myString, myString, myString)
	myString = "Синий"
	log.Println("Мой любимый цвет:", myString)
	fmt.Printf("Starting data: Adress: %p, Type: %T, Value: %#v\n", &myString, myString, myString)

	changeUserPointer(&myString)
	log.Println("Мой любимый цвет:", myString)
	fmt.Printf("Starting data: Adress: %p, Type: %T, Value: %#v\n", &myString, myString, myString)

}
