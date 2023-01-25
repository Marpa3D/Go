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

// Реализация функции взятия из стека. Pop
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func main() {
	var st Stack
	st.Push(10)
	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))

	st.Push(20)
	st.Push(30)
	st.Push(40)

	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))

	st.Pop()
	fmt.Printf("%+v\t adress: %p\t Длинна в байтах: %d\n", st, &st, len(st.items))
}
