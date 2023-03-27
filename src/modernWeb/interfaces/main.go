// Работа с интерфейсами
package main

import "log"

type Animal interface {
	Says() string
	NumberOfPets() int
}

type Cat struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name       string
	Color      string
	NumOfTeeth int
}

func main() {

	cat := Cat{
		Name:  "Bob",
		Breed: "Kamyshovaya",
	}
	PrintInfo(cat)

	gor := Gorilla{
		Name:       "BigFoot",
		Color:      "Black",
		NumOfTeeth: 34,
	}

	PrintInfo(gor)
}

func (c Cat) Says() string {
	return "Мяу!)"
}
func (c Cat) NumberOfPets() int {
	return 4
}

func (g Gorilla) Says() string {
	return "Арр!"
}
func (g Gorilla) NumberOfPets() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("Это животное говорит:", a.Says(), "У него лапок:", a.NumberOfPets())
}
