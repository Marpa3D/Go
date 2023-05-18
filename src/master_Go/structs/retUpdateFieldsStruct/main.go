// Извлечение и обновление полей структуры
package main

import "fmt"

type book struct {
	title  string
	author string
	year   int
}

func main() {
	newBook := book{title: "Грокаем алгориты"}
	fmt.Printf("\"%v\"\n", newBook.title)

	// Печать типа и полей - верба (глагол) #v
	fmt.Printf("%#v\n", newBook)

	// Печать полей структуры и их значений
	newBook.author = "Indi"
	newBook.year = 2013
	fmt.Printf("%+v", newBook)
}
