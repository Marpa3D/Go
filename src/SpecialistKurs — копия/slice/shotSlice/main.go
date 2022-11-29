// ДЗ. Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:
package main

import (
	"fmt"
)

// Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:
func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	res := []byte(text)
	s := '.'
	for i := 0; i <= width; i++ {

		res = res[:i]

	}
	if len(text) <= width {
		fmt.Println(text)
	} else {
		res = append(res, byte(s), byte(s), byte(s))
		fmt.Println(string(res))
	}

}
