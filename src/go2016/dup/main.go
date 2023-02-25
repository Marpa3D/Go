// реализация аналога uniq в UNIX системах
package main

import (
	"bufio"
	"fmt"
	"os"
)

//dup() - выводит из стандартного потока ввода повтор строки и колличество повторений
func dup() {
	counts := make(map[string]int) // хранилище входящих данных
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // чтение данных из os.Stdin
		counts[input.Text()]++
	}
	// итерация по карте
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
		}
	}

}
func main() {
	dup()
}
