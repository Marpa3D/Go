// практика по дженерикам
package main

import "fmt"

func main() {
	fmt.Println(Max(4, 8))

	var x, y uint = 5, 9

	fmt.Println(Max(x, y))
}

// Создаем свой constraint Ordered для работы логического x > y
type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func Max[T Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}
