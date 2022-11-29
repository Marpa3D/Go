// Практика с аргументами командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Аргументы: %s\n", args)
	fmt.Printf("Аргументы: %v", args[1:])
}
