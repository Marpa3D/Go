// Модульный тест для функции расчета числа Фибоначчи
package main

import (
	"testing"
)

func TestFibOne(t *testing.T) {
	actual := fibOne(8)
	if actual != 21 {
		t.Error("Ожидаемое значение 21 в позиции 8")
	}
}
func TestFibTwo(t *testing.T) {
	actual := fibTwo(8)
	if actual != 21 {
		t.Error("Ожидаемое значение 21 в позиции 8")
	}
}
