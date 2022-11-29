// Методы
package main

import "fmt"

// Прежде чем объявить метод, нужно создать структуру

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func main() {
	t := triangle{8}
	fmt.Println("Perimeter:", t.perimeter())
}
