// Пакет strings - вставка пробелов между символами строки
// Считаем овец перед сном...)))
package main

import (
	"fmt"
)

func main() {

	fmt.Println(countSheep(1))

}

func countSheep(num int) string {
	// Your code here!
	str := ""

	for i := 0; i <= num; i++ {
		str = fmt.Sprintf("%d sheep...", i)
		//str = string(i)
	}
	return str
}
