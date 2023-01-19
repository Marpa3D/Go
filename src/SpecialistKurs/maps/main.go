// Углубленное изучение хеш - карт, словарей и т.п., (map)
package main

import "fmt"

func main() {
	// Map - струтура данных в виде набора пар "ключ: значение".
	// 1. Объявление
	// var m map[string]string  Мапа - ссылочный тип, поэтому одного обьявления недостаточно!
	// m["foo"] = "bar"         Ошибка. Паника в runtime. panic: assignment to entry in nil map!

	// 1.1 Переменные типа map инициализируются с помощью функции make!
	// встроенная функция make - это конструктор для обьектов ссылочного типа.
	type myMap map[string]int
	var m myMap
	m = make(myMap, 5)
	m["bar"] = 44
	fmt.Printf("Значение: %v Тип: %Tn", m, m)

	// более компактный пример (внутри функции main)
	m1 := make(map[string]string)
	m1["Hi"] = "Mark"
	fmt.Printf("Значение: %v Тип: %Tn", m1, m1)
}
