// Практика с аргументами командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
}
