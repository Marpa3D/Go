// Два способа тестирования
package main

import "testing"

// Табличный тест
var tests = []struct {
	name     string
	divident float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Ожидали ошибку, но не получили")
			}
		} else {
			if err != nil {
				t.Error("Не ожидали ошибку, но получили ее", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("Ожидали %f, но получили %f", tt.expected, got)
		}
	}
}
