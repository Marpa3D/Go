// Пакет strings - вставка пробелов между символами строки
// Считаем овец перед сном...)))
package main

import (
	"fmt"
)

func main() {

	fmt.Println(countSheep(2))

}

func countSheep(num int) string {
	// Your code here!
	str := ""
	if num == 0 {
		fmt.Printf("0 sheep...")
	}
	for i := 1; i <= num; i++ {
		str += fmt.Sprintf("%d sheep...", i)
	}
	return str
}
