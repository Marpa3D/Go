//Программа, которая укорачивает строку до указанной длины и добавляет в конце многоточие
package main

import "fmt"

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	t := []byte(text)
	//t = t[:width]
	t = append(t[:width], '.', '.', '.')
	res := string(t)
	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	fmt.Println(res)
}
