// Работа с интерфейсами
package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

// Объвляем тип Book, который удовлетворяет интерфейсу fmt.Stringer
func (b Book) String() string {
	return fmt.Sprintf("Книга: %s - %s", b.Title, b.Author)
}

// Объвляем тип Count, который удовлетворяет интерфейсу fmt.Stringer
type Count int

// Объявляем функцию WriteLog(), которая берёт любой объект,
// удовлетворяющий интерфейсу fmt.Stringer в виде параметра.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func main() {
	//Инициализируем обьект Book и передаем его в WriteLog()
	book := Book{"Dune, part 1", "Frank Herbert"}
	WriteLog(book)

	//Инициализируем обьект Count и передаем его в WriteLog()
	count := Count(4)
	WriteLog(count)

	// В основной функции мы создали разные типы Book и Count, но передали их одной функции WriteLog().
	// А та вызвала соответствующие функции String() и записала результаты в журнал.
}
