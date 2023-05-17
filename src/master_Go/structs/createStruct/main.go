// Создание структуры
package main

import "fmt"

// Struct - это последовательность именнованых элементов, называемых полями.
// У каждого поля есть свое имя и тип данных. Полезны для гркппирования данных.
// Удобны для описания каких-либо сущностей, концепций (не только программых) этого мира.)
// Структура - это схема. Схема фиксируется во время компиляции (compile time)!!!
// Нельзя менять имена и типы данных полей структуры во время выполнения (runtime).

// Пример. Создадим структуру, описывающую книгу
type book struct {
	title  string
	author string
	year   int
}

func main() {
	// Создаем экземпляр струткутры типа book
	myBook := book{
		title:  "\"Dune\"",
		author: "F.Herbert",
		year:   1978,
	}
	fmt.Printf("Содержание переменной типа book: %+v\n\tТип переменной myBook: %T\n", myBook, myBook)
}
