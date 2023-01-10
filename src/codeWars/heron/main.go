// Функция heron, которая вычисляет площадь треугольника со сторонами a, b и c
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Heron(5.0, 2.0, 4.0))
}

func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	val := s * (s - a) * (s - b) * (s - c)
	area = math.Sqrt(val)
	return area
}
