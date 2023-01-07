// Пакет strings - вставка пробелов между символами строки
// Подсчет гласных во входящей строке
package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Hello WORLD!!!"
	v := strings.SplitAfter(s, "")
	fmt.Println(v)

	fmt.Println(GetCount("hello dolly and Jhon!"))

}

func GetCount(str string) (count int) {
	// Enter solution here
	count = 0
	for i, _ := range str {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			count++
		}
	}
	return count
}
