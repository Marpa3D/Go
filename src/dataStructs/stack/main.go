// Структуры данных. Стек
package main

import "fmt"

type Stack struct {
	items []int
}

// Реализация функции добавления в стек. Push
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func main() {
	var st Stack
	st.Push(10)
	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))

	st.Push(20)
	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))
}
