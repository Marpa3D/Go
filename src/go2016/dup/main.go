// реализация аналога uniq в UNIX системах
package main

import (
	"bufio"
	"fmt"
	"os"
)

//dup() - выводит из стандартного потока ввода повтор строки и колличество повторений
func dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, i := range counts {
		if i > 1 {
			fmt.Printf("%d\t%s", line, i)
		}
	}
}
func main() {
	dup()
}
