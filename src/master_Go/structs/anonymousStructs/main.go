// Анонимные структуры и поля
package main

import "fmt"

func main() {
	// 1. Анонимная структура - это структура, которая не имеет имени (явного псеводонима типа структуры).
	// Полезно, когда не нужно повторно использовать структуру в проекте!
	aStruct := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Mark",
		lastName:  "Evangelist",
		age:       98,
	}
	fmt.Printf("Поля анонимной струкутры: %#v\n", aStruct)
	fmt.Printf("Возраст: %#v лет.\n", aStruct.age)

	// 2. Анонимные поля структуры - поля без именб только типы данных
	type Book struct {
		string
		float64
		int
	}
	b := Book{"M.Plank", 22.8, 98}
	fmt.Printf("Values fields  struct variable b: %#v\n", b)

	// Вывод поля
	fmt.Printf("%v\n", b.string)

	// 3. Смешивание именованных и анонимных полей
	type Employee struct {
		name   string
		salary float64
		bool   // анонимное поле
	}
	e := Employee{"Jhon", 45000.0, false}
	fmt.Printf("%#v\n", e)

	e.bool = true
	fmt.Printf("%#v\n", e)

}
