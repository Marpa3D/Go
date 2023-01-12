//Программа, которая укорачивает строку до указанной длины и добавляет в конце многоточие
package main

import "fmt"

func main() {
	var text string
	var width int = 8
	fmt.Scanf("%s %d", &text, &width)

	if len(text) <= width {
		fmt.Println(text)
	} else {
		t := []byte(text)

		t = append(t[:width], '.', '.', '.')
		res := string(t)
		// Возьмите первые `width` байт строки `text`,
		// допишите в конце `...` и сохраните результат
		// в переменную `res`
		// ...

		fmt.Println(res)

	}
}
