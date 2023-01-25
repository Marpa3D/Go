// Функции для работы со стеком выведены в отдельный пакет
package stackFuncs

import "fmt"

type Stack struct {
	items []int
}

// Push  - реализация функции добавления в стек
func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

// Pop - реализация функции взятия из стека
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

// IsEmpty - проверка стека на наличие данных
func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

// Top возвращает последний добавленный елемент стека, если нет, то ноль и ошибку
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Стек пустой!")
	}
	return s.items[len(s.items)-1], nil
}

// Print - итерация по стеку и печать элементов в нем
func (s *Stack) Print() {
	for _, val := range s.items {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
