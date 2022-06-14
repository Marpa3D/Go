// Чтение из стандартного потока ввода, эхо - реакция
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	read := bufio.NewScanner(f)
	for read.Scan() {
		fmt.Println("=>", read.Text())
	}
}
