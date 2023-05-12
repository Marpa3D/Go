// Структуры данных. Реализация Set (множество, набор)
package main

import (
	"errors"
	"fmt"
)

// Set - наше описание заданной стркутуры данных
type Set struct {
	Elements map[string]struct{} // пустая структура для оптимизации памяти
}

// NewSet - конструктор экземпляра Set
func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// Add - добавление элемента в набор Set
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

// Delete - удаление элемента из Set, если он существует
func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("Элемент отсутствует в наборе Set")
	}
	delete(s.Elements, elem)
	return nil
}

// Contains - проверка на наличие элемента в наборе Set
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

// List - печатает список всех элементов в наборе Set
func (s *Set) List() {
	for key, _ := range s.Elements {
		fmt.Println(key)
	}
}

func main() {
	fmt.Println("Реализация структуры данных Set")
	mySet := NewSet()

	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Jupiter")
	mySet.Add("Mars")
	mySet.Add("Venus")
	mySet.Delete("Venus")

	mySet.List()
	fmt.Println(mySet.Contains("Earth"))
}
