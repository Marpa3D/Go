// Извлечение и обновление полей структуры
package main

import (
	"fmt"
	"strings"
)

type book struct {
	title  string
	author string
	year   int
}

func main() {
	myBook := book{title: "Грокаем алгориты"}
	fmt.Printf("\"%v\"\n", myBook.title)

	// Печать типа и полей - верба (глагол) #v
	fmt.Printf("%#v\n", myBook)

	// Печать полей структуры и их значений
	myBook.author = "Indi"
	myBook.year = 2013
	fmt.Printf("%+v\n", myBook)

	// Сравнение структур. Сравнивать можно (==, !=).
	// Структуры равны, если их соотвествующие поля равны!
	newBook := book{title: "Грокаем алгориты", author: "Indi", year: 2013}

	fmt.Println(strings.Repeat("*", 54))
	fmt.Println("Равны ли две структуры?)")
	fmt.Println(myBook == newBook)

	// Копирование структур
	fmt.Printf("Adress myBook: %p, Value: %+v\nAdress newBook: %p, Value: %+v\n", &myBook, myBook, &newBook, newBook)
	myBook = newBook
	myBook.year = 2018
	fmt.Printf("Adress myBook: %p, Value: %+v\nAdress newBook: %p, Value: %+v\n", &myBook, myBook, &newBook, newBook)
}
